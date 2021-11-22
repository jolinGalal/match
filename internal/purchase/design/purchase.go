package design

import (
	"github.com/jolinGalal/match/internal/purchase/types"
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("purchases", func() {
	dsl.Description("The events service exposes endpoints that require valid jwt token.")

	dsl.HTTP(func() {
		dsl.Path("/")
	})

	dsl.Error("unauthorized", func() {
		dsl.Description("the error returned when token is invalid")
	})
	dsl.Error("purchase-not-found", func() {
		dsl.Description("the error returned when purchase is not found")
	})

	dsl.Error("invalid-data", func() {
		dsl.Description("the error returned the purchase data is invalid")
	})

	dsl.Error("dublicate-data", func() {
		dsl.Description("the error returned the purchase data is already existing in the database")
	})
	dsl.Error("not-exist", func() {
		dsl.Description("the error returned when the purchase is not exists")
	})

	dsl.Method("deposit", func() {
		dsl.Description("Create new purchase")
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("deposit-purchase") // Enforce presence of "list-purchases" scope in JWT claims.
		})
		dsl.Payload(func() {

			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})
			dsl.Field(1, "id", dsl.Int64, func() {
				dsl.Description("the user id ")
			})
			dsl.Field(1, "deposit", dsl.Int, func() {
				dsl.Enum(types.Deposite[0], types.Deposite[1], types.Deposite[2], types.Deposite[3],
					types.Deposite[4])
			})
			dsl.Required("id", "deposit", "token")
		})

		dsl.HTTP(func() {
			dsl.POST("deposit/{id}")
			dsl.Response(dsl.StatusOK)
			dsl.Response("invalid-data", dsl.StatusNotAcceptable)
			dsl.Response("dublicate-data", dsl.StatusNotAcceptable)
		})
	})

	dsl.Method("buy", func() {
		dsl.Description("update existing purchase")
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("buy-purchase") // Enforce presence of "list-purchases" scope in JWT claims.
		})
		dsl.Payload(func() {
			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})
			dsl.Field(1, "id", dsl.Int64, func() {
				dsl.Description("the user id ")
			})
			dsl.Field(2, "purchase", dsl.ArrayOf(ProductsList), "")

			dsl.Required("id", "purchase", "token")
		})
		dsl.Result(PurchesList)
		dsl.HTTP(func() {
			dsl.POST("buy/{id}")
			dsl.Response(dsl.StatusOK)
			dsl.Response("invalid-data", dsl.StatusNotAcceptable)
			dsl.Response("dublicate-data", dsl.StatusNotAcceptable)
			dsl.Response("not-exist", dsl.StatusNotFound)
		})
	})
	dsl.Method("reset", func() {
		dsl.Description("reset deposit")
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("reset-purchase") // Enforce presence of "list-purchases" scope in JWT claims.
		})
		dsl.Payload(func() {

			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})
			dsl.Field(1, "id", dsl.Int64, func() {
				dsl.Description("the user id ")
			})

			dsl.Required("id", "token")
		})

		dsl.HTTP(func() {
			dsl.DELETE("reset/{id}")
			dsl.Response(dsl.StatusOK)
			dsl.Response("invalid-data", dsl.StatusNotAcceptable)
			dsl.Response("dublicate-data", dsl.StatusNotAcceptable)
		})
	})

})

// ProductsList ...
var ProductsList = dsl.Type("ProductsList", func() {
	dsl.Field(2, "productID", dsl.Int, "purchase name")
	dsl.Field(4, "purductAmount", dsl.Int, func() {
		dsl.Minimum(1)
	})
	dsl.Required("productID", "purductAmount")

})

var ProductInfo = dsl.Type("ProductInfo", func() {
	dsl.Field(2, "productName", dsl.String, "purchase name")
	dsl.Field(4, "purductAmount", dsl.Int, "purchase amount")
	dsl.Field(4, "purductUnitPrice", dsl.Int, "purchase price")
	dsl.Field(4, "purductTotalPrice", dsl.Int, "purchase price")
	dsl.Required("productName", "purductAmount", "purductTotalPrice", "purductUnitPrice")

})

var Changes = dsl.Type("Changes", func() {
	dsl.Field(2, "Coin", dsl.String, "Coin")
	dsl.Field(4, "Amount", dsl.Int, "amount")
	dsl.Required("Amount", "Coin")

})

// PurchesList ...
var PurchesList = dsl.Type("PurchesList", func() {
	dsl.Field(2, "ProductsList", dsl.ArrayOf(ProductInfo))

	dsl.Field(4, "Total", dsl.Int, "total amount")
	dsl.Field(5, "Changes", dsl.ArrayOf(Changes))
	dsl.Required("ProductsList", "Total", "Changes")

})
