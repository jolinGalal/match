// Code generated by goa v3.5.2, DO NOT EDIT.
//
// products HTTP client types
//
// Command:
// $ goa gen github.com/jolinGalal/match/internal/product/design

package client

import (
	products "github.com/jolinGalal/match/internal/product/gen/products"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "products" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	Product *ProductPayloadRequestBody `form:"product" json:"product" xml:"product"`
}

// UpdateRequestBody is the type of the "products" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Product *ProductPayloadRequestBody `form:"product" json:"product" xml:"product"`
}

// ListResponseBody is the type of the "products" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// events
	Products []*ListproductRespResponseBody `form:"products,omitempty" json:"products,omitempty" xml:"products,omitempty"`
	// pagination
	Pagination *PaginationResponseBody `form:"pagination,omitempty" json:"pagination,omitempty" xml:"pagination,omitempty"`
}

// ListUnauthorizedResponseBody is the type of the "products" service "list"
// endpoint HTTP response body for the "unauthorized" error.
type ListUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreateInvalidDataResponseBody is the type of the "products" service "create"
// endpoint HTTP response body for the "invalid-data" error.
type CreateInvalidDataResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreateDublicateDataResponseBody is the type of the "products" service
// "create" endpoint HTTP response body for the "dublicate-data" error.
type CreateDublicateDataResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateInvalidDataResponseBody is the type of the "products" service "update"
// endpoint HTTP response body for the "invalid-data" error.
type UpdateInvalidDataResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateDublicateDataResponseBody is the type of the "products" service
// "update" endpoint HTTP response body for the "dublicate-data" error.
type UpdateDublicateDataResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateNotExistResponseBody is the type of the "products" service "update"
// endpoint HTTP response body for the "not-exist" error.
type UpdateNotExistResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteNotExistResponseBody is the type of the "products" service "delete"
// endpoint HTTP response body for the "not-exist" error.
type DeleteNotExistResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ListproductRespResponseBody is used to define fields on response body types.
type ListproductRespResponseBody struct {
	// product Id
	ProductID *string `form:"productID,omitempty" json:"productID,omitempty" xml:"productID,omitempty"`
	// product name
	ProductName *string `form:"productName,omitempty" json:"productName,omitempty" xml:"productName,omitempty"`
	// product name
	ProductAmount *int `form:"productAmount,omitempty" json:"productAmount,omitempty" xml:"productAmount,omitempty"`
	// product role
	ProductCost *int `form:"productCost,omitempty" json:"productCost,omitempty" xml:"productCost,omitempty"`
	// product role
	SellerID *int `form:"sellerID,omitempty" json:"sellerID,omitempty" xml:"sellerID,omitempty"`
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
	ProductName   string `form:"productName" json:"productName" xml:"productName"`
	ProductAmount int    `form:"productAmount" json:"productAmount" xml:"productAmount"`
	ProductCost   int    `form:"productCost" json:"productCost" xml:"productCost"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "products" service.
func NewCreateRequestBody(p *products.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{}
	if p.Product != nil {
		body.Product = marshalProductsProductPayloadToProductPayloadRequestBody(p.Product)
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "products" service.
func NewUpdateRequestBody(p *products.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{}
	if p.Product != nil {
		body.Product = marshalProductsProductPayloadToProductPayloadRequestBody(p.Product)
	}
	return body
}

// NewListResultOK builds a "products" service "list" endpoint result from a
// HTTP "OK" response.
func NewListResultOK(body *ListResponseBody) *products.ListResult {
	v := &products.ListResult{}
	v.Products = make([]*products.ListproductResp, len(body.Products))
	for i, val := range body.Products {
		v.Products[i] = unmarshalListproductRespResponseBodyToProductsListproductResp(val)
	}
	v.Pagination = unmarshalPaginationResponseBodyToProductsPagination(body.Pagination)

	return v
}

// NewListUnauthorized builds a products service list endpoint unauthorized
// error.
func NewListUnauthorized(body *ListUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreateInvalidData builds a products service create endpoint invalid-data
// error.
func NewCreateInvalidData(body *CreateInvalidDataResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreateDublicateData builds a products service create endpoint
// dublicate-data error.
func NewCreateDublicateData(body *CreateDublicateDataResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateInvalidData builds a products service update endpoint invalid-data
// error.
func NewUpdateInvalidData(body *UpdateInvalidDataResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateDublicateData builds a products service update endpoint
// dublicate-data error.
func NewUpdateDublicateData(body *UpdateDublicateDataResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateNotExist builds a products service update endpoint not-exist error.
func NewUpdateNotExist(body *UpdateNotExistResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeleteNotExist builds a products service delete endpoint not-exist error.
func NewDeleteNotExist(body *DeleteNotExistResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateListResponseBody runs the validations defined on ListResponseBody
func ValidateListResponseBody(body *ListResponseBody) (err error) {
	if body.Products == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("products", "body"))
	}
	if body.Pagination == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pagination", "body"))
	}
	for _, e := range body.Products {
		if e != nil {
			if err2 := ValidateListproductRespResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateListUnauthorizedResponseBody runs the validations defined on
// list_unauthorized_response_body
func ValidateListUnauthorizedResponseBody(body *ListUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreateInvalidDataResponseBody runs the validations defined on
// create_invalid-data_response_body
func ValidateCreateInvalidDataResponseBody(body *CreateInvalidDataResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreateDublicateDataResponseBody runs the validations defined on
// create_dublicate-data_response_body
func ValidateCreateDublicateDataResponseBody(body *CreateDublicateDataResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateInvalidDataResponseBody runs the validations defined on
// update_invalid-data_response_body
func ValidateUpdateInvalidDataResponseBody(body *UpdateInvalidDataResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateDublicateDataResponseBody runs the validations defined on
// update_dublicate-data_response_body
func ValidateUpdateDublicateDataResponseBody(body *UpdateDublicateDataResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateNotExistResponseBody runs the validations defined on
// update_not-exist_response_body
func ValidateUpdateNotExistResponseBody(body *UpdateNotExistResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteNotExistResponseBody runs the validations defined on
// delete_not-exist_response_body
func ValidateDeleteNotExistResponseBody(body *DeleteNotExistResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateListproductRespResponseBody runs the validations defined on
// ListproductRespResponseBody
func ValidateListproductRespResponseBody(body *ListproductRespResponseBody) (err error) {
	if body.ProductID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productID", "body"))
	}
	if body.ProductName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productName", "body"))
	}
	if body.ProductAmount == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productAmount", "body"))
	}
	if body.ProductCost == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productCost", "body"))
	}
	if body.SellerID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sellerID", "body"))
	}
	return
}

// ValidateProductPayloadRequestBody runs the validations defined on
// productPayloadRequestBody
func ValidateProductPayloadRequestBody(body *ProductPayloadRequestBody) (err error) {
	if body.ProductAmount < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("body.productAmount", body.ProductAmount, 0, true))
	}
	if body.ProductCost < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("body.productCost", body.ProductCost, 0, true))
	}
	return
}
