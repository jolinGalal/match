package purchase

import (
	"context"
	"fmt"
	"log"

	"github.com/jolinGalal/match/internal/models"
	purchases "github.com/jolinGalal/match/internal/purchase/gen/purchases"
	"github.com/jolinGalal/match/internal/user/gen/users"
	"github.com/jolinGalal/match/pkg/auth"
	"github.com/jolinGalal/match/pkg/database"
	"goa.design/goa/v3/security"
)

// purchases service example implementation.
// The example methods log the requests and return zero values.
type purchasessrvc struct {
	logger *log.Logger
	auth   auth.AuthI
	db     database.GormI
}

var productI = models.NewProduct()
var userI = models.NewUser()

// NewPurchases returns the purchases service implementation.
func NewPurchases(logger *log.Logger, auth auth.AuthI, db database.GormI) purchases.Service {
	return &purchasessrvc{logger: logger, auth: auth, db: db}
}

// JWTAuth implements the authorization logic for service "purchases" for the
// "jwt" security scheme.
func (s *purchasessrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {

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
	if !s.auth.IsUser() {
		return ctx, users.MakeUnauthorized(fmt.Errorf("you are not a user"))
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
