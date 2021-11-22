package purchase

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jolinGalal/match/internal/models"
	purchases "github.com/jolinGalal/match/internal/purchase/gen/purchases"
)

// update existing purchase
func (s *purchasessrvc) Buy(ctx context.Context, p *purchases.BuyPayload) (res *purchases.PurchesList, err error) {
	res = &purchases.PurchesList{}
	s.getUser(p.ID)
	userObj, err := s.getUser(p.ID)
	if err != nil {
		return nil, err
	}
	products, err := s.getProducts(p.Purchase)
	if err != nil {
		return nil, err
	}
	err = getResponse(products, p.Purchase, res)
	if err != nil {
		return nil, err
	}
	if userObj.UserDeposit < res.Total {
		return nil, fmt.Errorf("you don't have enough money")
	}
	err = s.updatePrduct(products)
	if err != nil {
		return nil, err
	}
	s.getChanges(res, userObj)
	return
}

// getChanges ...
func (s *purchasessrvc) getChanges(res *purchases.PurchesList, userObj *models.User) {
	changes := userObj.UserDeposit - res.Total
	for d, _ := range models.Deposit {
		if changes < d {
			continue
		}
		c := &purchases.Changes{
			Coin:   strconv.Itoa(d),
			Amount: int(changes / d),
		}
		res.Changes = append(res.Changes, c)
		changes = changes - (d * c.Amount)

	}
	userObj.UserDeposit = changes
	s.db.Save(userObj)
}

// updatePrduct...
func (s *purchasessrvc) updatePrduct(p *[]models.Product) (err error) {
	s.db.Rest()
	for _, obj := range *p {
		fmt.Println(obj)
		err = s.db.Save(obj)
		if err != nil {
			return err
		}
	}
	return nil
}

// getUser
func (s *purchasessrvc) getUser(ID int64) (*models.User, error) {
	var userObj = &models.User{}
	s.db.Rest()
	err := s.db.Model(userObj).
		Where(fmt.Sprintf(" %s = %d", userI.ID(), ID)).
		Find(userObj)
	if userObj.UserID != ID {
		return nil, fmt.Errorf("user is not found")
	}
	return userObj, err

}

// getUser
func (s *purchasessrvc) getProducts(pur []*purchases.ProductsList) (*[]models.Product, error) {
	var productlst = []models.Product{}
	s.db.Rest()
	var query = []string{}
	for _, p := range pur {
		if p == nil {
			return nil, fmt.Errorf("product not found")
		}
		query = append(query, fmt.Sprintf("(%s=%d and %s>=%d)",
			productI.ID(), p.ProductID, productI.Amount(), p.PurductAmount))

	}
	err := s.db.Model(&models.Product{}).
		Where(strings.Join(query, " or ")).
		Find(&productlst)

	return &productlst, err

}

// getResponse ...
func getResponse(prod *[]models.Product, pur []*purchases.ProductsList,
	res *purchases.PurchesList) error {
	if res == nil {
		return fmt.Errorf("response not found")
	}
	if prod == nil {
		return fmt.Errorf("products not found")
	}
	var purchaseMap = ConvertPurchaseToMap(&pur)
	if len(purchaseMap) != len(*prod) {
		return fmt.Errorf("product item not found")
	}
	for _, obj := range *prod {
		m, _ := purchaseMap[int(obj.ProductID)]
		var p = &purchases.ProductInfo{
			ProductName:      obj.ProductName,
			PurductAmount:    m,
			PurductUnitPrice: obj.ProductCost,
		}
		p.PurductTotalPrice = m * obj.ProductCost
		obj.ProductAmount = obj.ProductAmount - m
		res.Total = res.Total + p.PurductTotalPrice
		res.ProductsList = append(res.ProductsList, p)
	}
	return nil
}

// ConvertPurchaseToMap ...
func ConvertPurchaseToMap(pur *[]*purchases.ProductsList) map[int]int {
	m := make(map[int]int)
	for _, obj := range *pur {
		if _, exists := m[obj.ProductID]; !exists {
			m[obj.ProductID] = obj.PurductAmount
		} else {
			m[obj.ProductID] = m[obj.ProductID] + obj.PurductAmount
		}

	}
	return m

}

// getTotal...
func getTotal(res *purchases.PurchesList) int {
	var sum = 0
	if res == nil {
		return sum
	}

	for _, obj := range res.ProductsList {
		sum += obj.PurductTotalPrice
	}
	return sum
}
