package product

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	products "github.com/jolinGalal/match/internal/product/gen/products"
	"github.com/jolinGalal/match/pkg/utils"
)

// update existing product
func (s *productssrvc) Update(ctx context.Context, p *products.UpdatePayload) (err error) {
	var productObj = &models.Product{}
	s.db.Rest()
	count, err := s.db.Model(productObj).
		Where(fmt.Sprintf(" lower(trim(%s)) like lower(trim('%s'))", productI.Name(), p.Product.ProductName)).
		Where(fmt.Sprintf(" %s != %d", productI.ID(), p.ID)).
		Count()
	if err != nil {
		return err
	}
	if *count > 0 {
		return fmt.Errorf("This product exists before")
	}
	// check the productname
	if err = utils.ValidString(p.Product.ProductName); err != nil {
		return err
	}

	// get the product
	s.db.Rest()
	err = s.db.Model(productObj).
		Where(fmt.Sprintf(" %s = %d", productI.ID(), p.ID)).
		Find(productObj)
	if err != nil {
		return nil
	}
	if productObj == nil {
		return fmt.Errorf("Id not found")
	}
	if productObj.ProductID == 0 {
		return fmt.Errorf("Id not found")
	}
	// update the product data
	productObj.ProductName = utils.PrepareString(&p.Product.ProductName, true)
	productObj.ProductAmount = p.Product.ProductAmount
	productObj.ProductCost = p.Product.ProductCost
	productObj.SellerID = int64(s.auth.GetUserID())
	s.db.Save(productObj)
	return
}
