package user

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	users "github.com/jolinGalal/match/internal/user/gen/users"
)

// login ...
func (s *userssrvc) Login(ctx context.Context, p *users.LoginPayload) (res *users.LoginRes, err error) {
	var userObj = &models.User{}
	s.db.Rest()
	err = s.db.Model(userObj).
		Where(fmt.Sprintf(" lower(trim(%s)) like lower(trim('%s'))", userI.Name(), p.UserName)).
		Where(fmt.Sprintf(" %s like '%s'", userI.Password(), p.USerPassword)).
		Find(userObj)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, users.MakeUserNotFound(
			fmt.Errorf("user is not found"),
		)
	}
	if userObj.UserName == "" {
		return nil, users.MakeUserNotFound(
			fmt.Errorf("user is not found"),
		)
	}

	s.auth, err = s.auth.UserID(userObj.UserID).Role(userObj.UserRole.String()).CreateToken()
	if err != nil {
		return nil, err
	}
	res = &users.LoginRes{
		Token: s.auth.GetToken(),
		ID:    int(userObj.UserID),
	}

	return res, err
}
