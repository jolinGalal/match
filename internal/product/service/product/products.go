package product

import (
	"context"
	"fmt"
	"log"

	"github.com/jolinGalal/match/internal/models"
	products "github.com/jolinGalal/match/internal/product/gen/products"
	"github.com/jolinGalal/match/internal/user/gen/users"
	"github.com/jolinGalal/match/pkg/auth"
	"github.com/jolinGalal/match/pkg/database"
	"goa.design/goa/v3/security"
)

var productI = models.NewProduct()
var userI = models.NewUser()

// products service example implementation.
// The example methods log the requests and return zero values.
type productssrvc struct {
	logger *log.Logger
	auth   auth.AuthI
	db     database.GormI
}

// NewProducts returns the products service implementation.
func NewProducts(logger *log.Logger, auth auth.AuthI, db database.GormI) products.Service {
	return &productssrvc{logger: logger, auth: auth, db: db}
}

// JWTAuth implements the authorization logic for service "products" for the
// "jwt" security scheme.
func (s *productssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	// check the auth token
	if token == "" {
		return ctx, users.MakeUnauthorized(
			fmt.Errorf("missing authentication header"),
		)
	}
	var err error
	s.auth, err = s.auth.VerifyToken(token)
	if err != nil {
		return ctx, users.MakeUnauthorized(
			err,
		)
	}
	if len(scheme.RequiredScopes) > 0 {
		if scheme.RequiredScopes[0] == "create-product" ||
			scheme.RequiredScopes[0] == "delete-product" ||
			scheme.RequiredScopes[0] == "update-product" {
			if !s.auth.IsAdmin() {
				return ctx, users.MakeUnauthorized(fmt.Errorf("you are not an admin"))
			}

		}
	}
	s.db.Rest()
	var userObj = &models.User{}
	count, err := s.db.Model(userObj).
		Where(fmt.Sprintf(" %s = %d", userI.ID(), s.auth.GetUserID())).
		Count()
	if err != nil {
		return ctx, err
	}
	if *count == 0 {
		return ctx, fmt.Errorf("This user is not exists")
	}

	return ctx, nil
}
