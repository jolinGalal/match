package user

import (
	"context"
	"fmt"
	"log"

	"github.com/jolinGalal/match/internal/models"
	users "github.com/jolinGalal/match/internal/user/gen/users"
	"github.com/jolinGalal/match/pkg/auth"
	"github.com/jolinGalal/match/pkg/database"
	"goa.design/goa/v3/security"
)

var userI = models.NewUser()

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
	auth   auth.AuthI
	db     database.GormI
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger, auth auth.AuthI, db database.GormI) users.Service {
	return &userssrvc{logger: logger, auth: auth, db: db}
}

// JWTAuth implements the authorization logic for service "users" for the "jwt"
// security scheme.
func (s *userssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {

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
