// Code generated by goa v3.5.2, DO NOT EDIT.
//
// match HTTP client CLI support package
//
// Command:
// $ goa gen github.com/jolinGalal/match/internal/user/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	usersc "github.com/jolinGalal/match/internal/user/gen/http/users/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `users (list|create|update|delete|login)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` users list --sort-direction "asc" --sort-key "UserName" --page-number 476080364796613033 --page-size 578784000480783521 --token "Nulla mollitia at magni ea amet earum."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		usersFlags = flag.NewFlagSet("users", flag.ContinueOnError)

		usersListFlags             = flag.NewFlagSet("list", flag.ExitOnError)
		usersListSortDirectionFlag = usersListFlags.String("sort-direction", "desc", "")
		usersListSortKeyFlag       = usersListFlags.String("sort-key", "CreatedAt", "")
		usersListPageNumberFlag    = usersListFlags.String("page-number", "1", "")
		usersListPageSizeFlag      = usersListFlags.String("page-size", "20", "")
		usersListTokenFlag         = usersListFlags.String("token", "REQUIRED", "")

		usersCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		usersCreateBodyFlag = usersCreateFlags.String("body", "REQUIRED", "")

		usersUpdateFlags     = flag.NewFlagSet("update", flag.ExitOnError)
		usersUpdateBodyFlag  = usersUpdateFlags.String("body", "REQUIRED", "")
		usersUpdateIDFlag    = usersUpdateFlags.String("id", "REQUIRED", "the user id to be updated")
		usersUpdateTokenFlag = usersUpdateFlags.String("token", "REQUIRED", "")

		usersDeleteFlags     = flag.NewFlagSet("delete", flag.ExitOnError)
		usersDeleteIDFlag    = usersDeleteFlags.String("id", "REQUIRED", "the user id to be updated")
		usersDeleteTokenFlag = usersDeleteFlags.String("token", "REQUIRED", "")

		usersLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		usersLoginBodyFlag = usersLoginFlags.String("body", "REQUIRED", "")
	)
	usersFlags.Usage = usersUsage
	usersListFlags.Usage = usersListUsage
	usersCreateFlags.Usage = usersCreateUsage
	usersUpdateFlags.Usage = usersUpdateUsage
	usersDeleteFlags.Usage = usersDeleteUsage
	usersLoginFlags.Usage = usersLoginUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "users":
			svcf = usersFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "users":
			switch epn {
			case "list":
				epf = usersListFlags

			case "create":
				epf = usersCreateFlags

			case "update":
				epf = usersUpdateFlags

			case "delete":
				epf = usersDeleteFlags

			case "login":
				epf = usersLoginFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "users":
			c := usersc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = usersc.BuildListPayload(*usersListSortDirectionFlag, *usersListSortKeyFlag, *usersListPageNumberFlag, *usersListPageSizeFlag, *usersListTokenFlag)
			case "create":
				endpoint = c.Create()
				data, err = usersc.BuildCreatePayload(*usersCreateBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = usersc.BuildUpdatePayload(*usersUpdateBodyFlag, *usersUpdateIDFlag, *usersUpdateTokenFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = usersc.BuildDeletePayload(*usersDeleteIDFlag, *usersDeleteTokenFlag)
			case "login":
				endpoint = c.Login()
				data, err = usersc.BuildLoginPayload(*usersLoginBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// usersUsage displays the usage of the users command and its subcommands.
func usersUsage() {
	fmt.Fprintf(os.Stderr, `The events service exposes endpoints that require valid jwt token.
Usage:
    %[1]s [globalflags] users COMMAND [flags]

COMMAND:
    list: list users
    create: Create new user
    update: update existing user
    delete: delete existing user
    login: delete existing user

Additional help:
    %[1]s users COMMAND --help
`, os.Args[0])
}
func usersListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] users list -sort-direction STRING -sort-key STRING -page-number INT -page-size INT -token STRING

list users
    -sort-direction STRING: 
    -sort-key STRING: 
    -page-number INT: 
    -page-size INT: 
    -token STRING: 

Example:
    %[1]s users list --sort-direction "asc" --sort-key "UserName" --page-number 476080364796613033 --page-size 578784000480783521 --token "Nulla mollitia at magni ea amet earum."
`, os.Args[0])
}

func usersCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] users create -body JSON

Create new user
    -body JSON: 

Example:
    %[1]s users create --body '{
      "user": {
         "UserName": "Est sed in.",
         "UserPassword": "Voluptates ut provident in eos sit quo.",
         "UserRole": "seller"
      }
   }'
`, os.Args[0])
}

func usersUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] users update -body JSON -id INT64 -token STRING

update existing user
    -body JSON: 
    -id INT64: the user id to be updated
    -token STRING: 

Example:
    %[1]s users update --body '{
      "user": {
         "UserName": "Est sed in.",
         "UserPassword": "Voluptates ut provident in eos sit quo.",
         "UserRole": "seller"
      }
   }' --id 8680840607413305103 --token "Corporis iusto perspiciatis et qui."
`, os.Args[0])
}

func usersDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] users delete -id INT64 -token STRING

delete existing user
    -id INT64: the user id to be updated
    -token STRING: 

Example:
    %[1]s users delete --id 1261404873507736171 --token "Fugit qui."
`, os.Args[0])
}

func usersLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] users login -body JSON

delete existing user
    -body JSON: 

Example:
    %[1]s users login --body '{
      "USerPassword": "Expedita qui velit eveniet atque.",
      "UserName": "Quia rem debitis corporis esse."
   }'
`, os.Args[0])
}