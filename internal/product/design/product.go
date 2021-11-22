package design

import (
	"github.com/jolinGalal/match/internal/product/types"
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("products", func() {
	dsl.Description("The events service exposes endpoints that require valid jwt token.")

	dsl.HTTP(func() {
		dsl.Path("/")
	})

	dsl.Error("unauthorized", func() {
		dsl.Description("the error returned when token is invalid")
	})
	dsl.Error("product-not-found", func() {
		dsl.Description("the error returned when product is not found")
	})

	dsl.Error("invalid-data", func() {
		dsl.Description("the error returned the product data is invalid")
	})

	dsl.Error("dublicate-data", func() {
		dsl.Description("the error returned the product data is already existing in the database")
	})
	dsl.Error("not-exist", func() {
		dsl.Description("the error returned when the product is not exists")
	})

	dsl.Method("list", func() {
		dsl.Description("list products")

		// The list products endpoint is secured via jwt auth
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("list-products") // Enforce presence of "list-products" scope in JWT claims.
		})

		dsl.Payload(func() {
			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})

			dsl.Field(5, "sort_key", dsl.String, func() {
				dsl.Description("")
				dsl.Enum(types.ProductSort.CreatedAt, types.ProductSort.ProductID, types.ProductSort.ProductName)
				dsl.Default(types.ProductSort.CreatedAt)
			})
			dsl.Field(6, "sort_direction", dsl.String, func() {
				dsl.Enum(types.SortDirection.Asc, types.SortDirection.Desc)
				dsl.Default(types.SortDirection.Desc)
			})
			dsl.Field(7, "page_number", dsl.Int, func() {
				dsl.Default(1)
			})
			dsl.Field(8, "page_size", dsl.Int, func() {
				dsl.Default(20)
			})
			dsl.Required("token")
		})

		dsl.Result(func() {
			dsl.Field(1, "products", dsl.ArrayOf(ListproductResp), "events")
			dsl.Field(2, "pagination", Pagination, "pagination")
			dsl.Required("products", "pagination")
		})

		dsl.HTTP(func() {
			dsl.GET("")
			dsl.Params(func() {
				dsl.Param("sort_direction")
				dsl.Param("sort_key")
				dsl.Param("page_number")
				dsl.Param("page_size")

			})
			dsl.Response(dsl.StatusOK)
			dsl.Response("unauthorized", dsl.StatusUnauthorized)
		})

	})

	dsl.Method("create", func() {
		dsl.Description("Create new product")
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("create-product") // Enforce presence of "list-products" scope in JWT claims.
		})
		dsl.Payload(func() {

			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})
			dsl.Field(1, "product", productPayload, "")

			dsl.Required("product", "token")
		})

		dsl.HTTP(func() {
			dsl.POST("")
			dsl.Response(dsl.StatusOK)
			dsl.Response("invalid-data", dsl.StatusNotAcceptable)
			dsl.Response("dublicate-data", dsl.StatusNotAcceptable)
		})
	})

	dsl.Method("update", func() {
		dsl.Description("update existing product")
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("update-product") // Enforce presence of "list-products" scope in JWT claims.
		})
		dsl.Payload(func() {
			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})
			dsl.Field(1, "id", dsl.Int64, func() {
				dsl.Description("the product id to be updated")
			})
			dsl.Field(2, "product", productPayload, "")

			dsl.Required("id", "product", "token")
		})

		dsl.HTTP(func() {
			dsl.PUT("/{id}")
			dsl.Response(dsl.StatusOK)
			dsl.Response("invalid-data", dsl.StatusNotAcceptable)
			dsl.Response("dublicate-data", dsl.StatusNotAcceptable)
			dsl.Response("not-exist", dsl.StatusNotFound)
		})
	})

	dsl.Method("delete", func() {
		dsl.Description("delete existing product")
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("delete-product") // Enforce presence of "list-products" scope in JWT claims.
		})
		dsl.Payload(func() {
			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})
			dsl.Field(2, "id", dsl.Int64, func() {
				dsl.Description("the product id to be updated")
			})
			dsl.Required("token", "id")
		})

		dsl.HTTP(func() {
			dsl.DELETE("/{id}")
			dsl.Response(dsl.StatusOK)
			dsl.Response("not-exist", dsl.StatusNotFound)
		})
	})

})

// ListproductResp ...
var ListproductResp = dsl.Type("ListproductResp", func() {
	dsl.Field(1, "productID", dsl.String, "product Id")
	dsl.Field(2, "productName", dsl.String, "product name")
	dsl.Field(4, "productAmount", dsl.Int, "product name")
	dsl.Field(5, "productCost", dsl.Int, "product role")
	dsl.Field(6, "sellerID", dsl.Int, "product role")
	dsl.Required("productID", "productName", "productAmount", "productCost", "sellerID")

})

// productPayload ...
var productPayload = dsl.Type("productPayload", func() {
	dsl.Field(1, "productName", dsl.String, "product name")
	dsl.Field(2, "productAmount", dsl.Int, func() {
		dsl.Minimum(0)
	})
	dsl.Field(4, "productCost", dsl.Int, func() {
		dsl.Minimum(0)
	})
	dsl.Required("productName", "productAmount", "productCost")
})

// Pagination is
var Pagination = dsl.ResultType("Pagination", func() {
	dsl.Field(1, "current_page", dsl.Int, func() {
		dsl.Description("The current page")
		dsl.Example(1)
	})
	dsl.Field(1, "page_size", dsl.Int, func() {
		dsl.Description("Max number of records per page ")
		dsl.Example(3)
	})
	dsl.Field(1, "total_pages", dsl.Int, func() {
		dsl.Description("Total pages")
		dsl.Example(11)
	})
	dsl.Field(1, "total_count", dsl.Int64, func() {
		dsl.Description("Total records count")
		dsl.Example(33)
	})
})
