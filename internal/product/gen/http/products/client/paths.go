// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the products service.
//
// Command:
// $ goa gen github.com/jolinGalal/match/internal/product/design

package client

import (
	"fmt"
)

// ListProductsPath returns the URL path to the products service list HTTP endpoint.
func ListProductsPath() string {
	return "/product/v1"
}

// CreateProductsPath returns the URL path to the products service create HTTP endpoint.
func CreateProductsPath() string {
	return "/product/v1"
}

// UpdateProductsPath returns the URL path to the products service update HTTP endpoint.
func UpdateProductsPath(id int64) string {
	return fmt.Sprintf("/product/v1/%v", id)
}

// DeleteProductsPath returns the URL path to the products service delete HTTP endpoint.
func DeleteProductsPath(id int64) string {
	return fmt.Sprintf("/product/v1/%v", id)
}
