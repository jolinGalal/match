package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jolinGalal/match/internal/models"
	products "github.com/jolinGalal/match/internal/product/gen/products"
	product_service "github.com/jolinGalal/match/internal/product/service/product"
	"github.com/jolinGalal/match/pkg/auth"
	"github.com/jolinGalal/match/pkg/database"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF       = flag.String("host", "localhost", "Server host (valid values: localhost)")
		hostPort    = flag.String("host_port", "8099", "Server host port (valid values: 8090)")
		domainF     = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF   = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF     = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF        = flag.Bool("debug", false, "Log request and response bodies")
		dbUser      = flag.String("db_user", "postgres", "")
		dbPassword  = flag.String("db_password", "example", "")
		dbHost      = flag.String("db_host", "localhost", "")
		dbName      = flag.String("db_name", "match", "")
		dbPort      = flag.Int("db_port", 5432, "")
		dbSSLMode   = flag.String("db_ssl_mode", "disable", "database SSL mode")
		dbMigrate   = flag.Bool("db_Migrate", true, "")
		dbCreate    = flag.Bool("db_Create", true, "")
		jwtToken    = flag.String("jwt_token", "jwttokenkeys", "")
		jwtInterval = flag.Int("jwt_interval", 300, "token expirtion time")
	)
	flag.Parse()

	// create database
	db, err := database.New().Host(*dbHost).DbName(*dbName).UserName(*dbUser).Password(*dbPassword).
		Sslmode(*dbSSLMode).Port(*dbPort).Migrate(*dbMigrate).Create(*dbCreate).
		Models(&models.User{}).Models(&models.Product{}).Get()
	if err != nil {
		log.Fatal("error in creation db", err)
	}
	dbgorm := database.NewDB(db)
	// set up jwt global key
	os.Setenv("JWT_ACCESS_SECRET", *jwtToken)

	// create authentication
	auth, err := auth.NewAuth()
	if err != nil {
		log.Fatal("error in authentication ", err)
	}
	auth = auth.Interval(*jwtInterval)

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[match] ", log.Ltime)
	}

	// Initialize the services.
	var (
		productsSvc products.Service
	)
	{
		productsSvc = product_service.NewProducts(logger, auth, dbgorm)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		productsEndpoints *products.Endpoints
	)
	{
		productsEndpoints = products.NewEndpoints(productsSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:" + *hostPort
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", u.Host, err)
					os.Exit(1)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, productsEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
