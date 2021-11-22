package models

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

// Product model
type Product struct {
	gorm.Model
	ProductID     int64 `gorm:"primaryKey;unique;autoIncrement:true" `
	ProductName   string
	ProductAmount int
	ProductCost   int
	SellerID      int64 `gorm:"foreignKey:user_id;references:product_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
}

// ProductInterface ...
type ProductInterface interface {
	GetTableName() string
	ID() string
	Name() string
	Amount() string
	Cost() string
	Seller() string
}

// productLock ...
var productLock = &sync.Mutex{}

// productInstance ...
var productInstance *Product

// NewProduct ...
func NewProduct() ProductInterface {
	if productInstance == nil {
		productLock.Lock()
		defer productLock.Unlock()
		if productInstance == nil {
			productInstance = &Product{}
		}
	}

	return productInstance
}

// GetTableName ...
func (p *Product) GetTableName() string {
	return "products"
}

// ID ...
func (p *Product) ID() string {
	return fmt.Sprintf("%s.product_id", p.GetTableName())
}

// Name ...
func (p *Product) Name() string {
	return fmt.Sprintf("%s.product_name", p.GetTableName())
}

// Amount ...
func (p *Product) Amount() string {
	return fmt.Sprintf("%s.product_amount", p.GetTableName())
}

// Cost ...
func (p *Product) Cost() string {
	return fmt.Sprintf("%s.product_cost", p.GetTableName())
}

// Seller ...
func (p *Product) Seller() string {
	return fmt.Sprintf("%s.seller_id", p.GetTableName())
}

//  CreatedAt...
func (p *Product) CreatedAt() string {
	return fmt.Sprintf("%s.created_at", p.GetTableName())
}
