package product

import (
	"context"
	"fmt"

	"github.com/jolinGalal/match/internal/models"
	products "github.com/jolinGalal/match/internal/product/gen/products"
)

// delete existing product
func (s *productssrvc) Delete(ctx context.Context, p *products.DeletePayload) (err error) {
	var productObj = &models.Product{}

	s.db.Rest()
	count, err := s.db.Model(productObj).
		Where(fmt.Sprintf(" %s = %d", productI.ID(), p.ID)).
		Count()
	if err != nil {
		return err
	}
	if *count == 0 {
		return fmt.Errorf("This product is not existing")
	}
	fmt.Println(productObj)

	err = s.db.Delete(productObj, p.ID)

	return err
}
