// Code generated by goa v3.5.2, DO NOT EDIT.
//
// products HTTP server types
//
// Command:
// $ goa gen github.com/jolinGalal/match/internal/product/design

package server

import (
	products "github.com/jolinGalal/match/internal/product/gen/products"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "products" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	Product *ProductPayloadRequestBody `form:"product,omitempty" json:"product,omitempty" xml:"product,omitempty"`
}

// UpdateRequestBody is the type of the "products" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Product *ProductPayloadRequestBody `form:"product,omitempty" json:"product,omitempty" xml:"product,omitempty"`
}

// ListResponseBody is the type of the "products" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// events
	Products []*ListproductRespResponseBody `form:"products" json:"products" xml:"products"`
	// pagination
	Pagination *PaginationResponseBody `form:"pagination" json:"pagination" xml:"pagination"`
}

// ListUnauthorizedResponseBody is the type of the "products" service "list"
// endpoint HTTP response body for the "unauthorized" error.
type ListUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateInvalidDataResponseBody is the type of the "products" service "create"
// endpoint HTTP response body for the "invalid-data" error.
type CreateInvalidDataResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateDublicateDataResponseBody is the type of the "products" service
// "create" endpoint HTTP response body for the "dublicate-data" error.
type CreateDublicateDataResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInvalidDataResponseBody is the type of the "products" service "update"
// endpoint HTTP response body for the "invalid-data" error.
type UpdateInvalidDataResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateDublicateDataResponseBody is the type of the "products" service
// "update" endpoint HTTP response body for the "dublicate-data" error.
type UpdateDublicateDataResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateNotExistResponseBody is the type of the "products" service "update"
// endpoint HTTP response body for the "not-exist" error.
type UpdateNotExistResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteNotExistResponseBody is the type of the "products" service "delete"
// endpoint HTTP response body for the "not-exist" error.
type DeleteNotExistResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListproductRespResponseBody is used to define fields on response body types.
type ListproductRespResponseBody struct {
	// product Id
	ProductID string `form:"productID" json:"productID" xml:"productID"`
	// product name
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// product name
	ProductAmount int `form:"productAmount" json:"productAmount" xml:"productAmount"`
	// product role
	ProductCost int `form:"productCost" json:"productCost" xml:"productCost"`
	// product role
	SellerID int `form:"sellerID" json:"sellerID" xml:"sellerID"`
}

// PaginationResponseBody is used to define fields on response body types.
type PaginationResponseBody struct {
	// The current page
	CurrentPage *int `form:"current_page,omitempty" json:"current_page,omitempty" xml:"current_page,omitempty"`
	// Max number of records per page
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Total pages
	TotalPages *int `form:"total_pages,omitempty" json:"total_pages,omitempty" xml:"total_pages,omitempty"`
	// Total records count
	TotalCount *int64 `form:"total_count,omitempty" json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// ProductPayloadRequestBody is used to define fields on request body types.
type ProductPayloadRequestBody struct {
	// product name
	ProductName   *string `form:"productName,omitempty" json:"productName,omitempty" xml:"productName,omitempty"`
	ProductAmount *int    `form:"productAmount,omitempty" json:"productAmount,omitempty" xml:"productAmount,omitempty"`
	ProductCost   *int    `form:"productCost,omitempty" json:"productCost,omitempty" xml:"productCost,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "products" service.
func NewListResponseBody(res *products.ListResult) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Products != nil {
		body.Products = make([]*ListproductRespResponseBody, len(res.Products))
		for i, val := range res.Products {
			body.Products[i] = marshalProductsListproductRespToListproductRespResponseBody(val)
		}
	}
	if res.Pagination != nil {
		body.Pagination = marshalProductsPaginationToPaginationResponseBody(res.Pagination)
	}
	return body
}

// NewListUnauthorizedResponseBody builds the HTTP response body from the
// result of the "list" endpoint of the "products" service.
func NewListUnauthorizedResponseBody(res *goa.ServiceError) *ListUnauthorizedResponseBody {
	body := &ListUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateInvalidDataResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "products" service.
func NewCreateInvalidDataResponseBody(res *goa.ServiceError) *CreateInvalidDataResponseBody {
	body := &CreateInvalidDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateDublicateDataResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "products" service.
func NewCreateDublicateDataResponseBody(res *goa.ServiceError) *CreateDublicateDataResponseBody {
	body := &CreateDublicateDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInvalidDataResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "products" service.
func NewUpdateInvalidDataResponseBody(res *goa.ServiceError) *UpdateInvalidDataResponseBody {
	body := &UpdateInvalidDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateDublicateDataResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "products" service.
func NewUpdateDublicateDataResponseBody(res *goa.ServiceError) *UpdateDublicateDataResponseBody {
	body := &UpdateDublicateDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateNotExistResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "products" service.
func NewUpdateNotExistResponseBody(res *goa.ServiceError) *UpdateNotExistResponseBody {
	body := &UpdateNotExistResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteNotExistResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "products" service.
func NewDeleteNotExistResponseBody(res *goa.ServiceError) *DeleteNotExistResponseBody {
	body := &DeleteNotExistResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListPayload builds a products service list endpoint payload.
func NewListPayload(sortDirection string, sortKey string, pageNumber int, pageSize int, token string) *products.ListPayload {
	v := &products.ListPayload{}
	v.SortDirection = sortDirection
	v.SortKey = sortKey
	v.PageNumber = pageNumber
	v.PageSize = pageSize
	v.Token = token

	return v
}

// NewCreatePayload builds a products service create endpoint payload.
func NewCreatePayload(body *CreateRequestBody, token string) *products.CreatePayload {
	v := &products.CreatePayload{}
	v.Product = unmarshalProductPayloadRequestBodyToProductsProductPayload(body.Product)
	v.Token = token

	return v
}

// NewUpdatePayload builds a products service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id int64, token string) *products.UpdatePayload {
	v := &products.UpdatePayload{}
	v.Product = unmarshalProductPayloadRequestBodyToProductsProductPayload(body.Product)
	v.ID = id
	v.Token = token

	return v
}

// NewDeletePayload builds a products service delete endpoint payload.
func NewDeletePayload(id int64, token string) *products.DeletePayload {
	v := &products.DeletePayload{}
	v.ID = id
	v.Token = token

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Product == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("product", "body"))
	}
	if body.Product != nil {
		if err2 := ValidateProductPayloadRequestBody(body.Product); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Product == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("product", "body"))
	}
	if body.Product != nil {
		if err2 := ValidateProductPayloadRequestBody(body.Product); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProductPayloadRequestBody runs the validations defined on
// productPayloadRequestBody
func ValidateProductPayloadRequestBody(body *ProductPayloadRequestBody) (err error) {
	if body.ProductName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productName", "body"))
	}
	if body.ProductAmount == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productAmount", "body"))
	}
	if body.ProductCost == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productCost", "body"))
	}
	if body.ProductAmount != nil {
		if *body.ProductAmount < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.productAmount", *body.ProductAmount, 0, true))
		}
	}
	if body.ProductCost != nil {
		if *body.ProductCost < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.productCost", *body.ProductCost, 0, true))
		}
	}
	return
}
