package purchase

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	purchases "github.com/jolinGalal/match/internal/purchase/gen/purchases"
)

// Create new purchase
func (s *purchasessrvc) Deposit(ctx context.Context, p *purchases.DepositPayload) (err error) {
	var userObj = &models.User{}
	// check if the user is the same as the login one
	if p.ID != int64(s.auth.GetUserID()) {
		return fmt.Errorf("invalid credential")
	}
	s.db.Rest()
	// get the user
	err = s.db.Model(userObj).
		Where(fmt.Sprintf(" %s = %d", userI.ID(), p.ID)).
		Find(userObj)
	if err != nil {
		return nil
	}
	if userObj == nil {
		return fmt.Errorf("Id not found")
	}
	if userObj.UserID == 0 {
		return fmt.Errorf("Id not found")
	}
	var query = fmt.Sprintf("update %s set user_deposit=%d where %s=%d", userI.GetTableName(),
		userObj.UserDeposit+p.Deposit, userI.ID(), userObj.UserID)

	return s.db.Execute(query)
}
