package product

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	products "github.com/jolinGalal/match/internal/product/gen/products"
	"github.com/jolinGalal/match/pkg/utils"
)

// Create new product
func (s *productssrvc) Create(ctx context.Context, p *products.CreatePayload) (err error) {
	var productObj = &models.Product{}
	// check if the product exists before
	s.db.Rest()
	count, err := s.db.Model(productObj).
		Where(fmt.Sprintf(" lower(trim(%s)) like lower(trim('%s'))", productI.Name(), p.Product.ProductName)).
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
	// create new product
	productObj = &models.Product{
		ProductName:   utils.PrepareString(&p.Product.ProductName, true),
		ProductAmount: p.Product.ProductAmount,
		ProductCost:   p.Product.ProductCost,
		SellerID:      int64(s.auth.GetUserID()),
	}

	s.db.Save(productObj)

	return
}
