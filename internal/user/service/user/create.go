package user

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	users "github.com/jolinGalal/match/internal/user/gen/users"
	"github.com/jolinGalal/match/pkg/utils"
)

// Create new user
func (s *userssrvc) Create(ctx context.Context, p *users.CreatePayload) (err error) {
	var userObj = &models.User{}
	// check if the user exists before
	s.db.Rest()
	count, err := s.db.Model(userObj).
		Where(fmt.Sprintf(" lower(trim(%s)) like lower(trim('%s'))", userI.Name(), p.User.UserName)).
		Count()
	if err != nil {
		return err
	}
	if *count > 0 {
		return fmt.Errorf("This user exists before")
	}
	// check the username
	if err = utils.ValidString(p.User.UserName); err != nil {
		return err
	}
	// check the password
	if err = utils.ValidPassword(p.User.UserPassword); err != nil {
		return err
	}
	// create new user
	userObj = &models.User{
		UserName:     utils.PrepareString(&p.User.UserName, true),
		UserPassword: utils.ConvertToSha(p.User.UserPassword),
	}
	userObj.UserRole, err = userI.ConvertStringToRole(p.User.UserRole)
	if err != nil {
		return err
	}

	// if p.User.UserDeposit != nil {
	// 	if !userI.ValidDeposit(*p.User.UserDeposit) {
	// 		return fmt.Errorf("Invalid deposit")
	// 	}
	// 	userObj.UserDeposit = *p.User.UserDeposit

	// } else {
	// 	userObj.UserDeposit = 0
	// }
	s.db.Save(userObj)

	return
}
