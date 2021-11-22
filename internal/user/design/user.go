package design

import (
	"github.com/jolinGalal/match/internal/models"
	"github.com/jolinGalal/match/internal/user/types"
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("users", func() {
	dsl.Description("The events service exposes endpoints that require valid jwt token.")

	dsl.HTTP(func() {
		dsl.Path("/")
	})

	dsl.Error("unauthorized", func() {
		dsl.Description("the error returned when token is invalid")
	})
	dsl.Error("user-not-found", func() {
		dsl.Description("the error returned when user is not found")
	})

	dsl.Error("invalid-data", func() {
		dsl.Description("the error returned the user data is invalid")
	})

	dsl.Error("dublicate-data", func() {
		dsl.Description("the error returned the user data is already existing in the database")
	})
	dsl.Error("not-exist", func() {
		dsl.Description("the error returned when the user is not exists")
	})

	dsl.Method("list", func() {
		dsl.Description("list users")

		// The list users endpoint is secured via jwt auth
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("list-users") // Enforce presence of "list-users" scope in JWT claims.
		})

		dsl.Payload(func() {
			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})

			dsl.Field(5, "sort_key", dsl.String, func() {
				dsl.Description("")
				dsl.Enum(types.UserSort.CreatedAt, types.UserSort.UserID, types.UserSort.UserName)
				dsl.Default(types.UserSort.CreatedAt)
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
			dsl.Field(1, "users", dsl.ArrayOf(ListUserResp), "events")
			dsl.Field(2, "pagination", Pagination, "pagination")
			dsl.Required("users", "pagination")
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
		dsl.Description("Create new user")
		dsl.Payload(func() {

			dsl.Field(1, "user", UserPayload, "")

			dsl.Required("user")
		})

		dsl.HTTP(func() {
			dsl.POST("")
			dsl.Response(dsl.StatusOK)
			dsl.Response("invalid-data", dsl.StatusNotAcceptable)
			dsl.Response("dublicate-data", dsl.StatusNotAcceptable)
		})
	})

	dsl.Method("update", func() {
		dsl.Description("update existing user")
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("update-user") // Enforce presence of "list-users" scope in JWT claims.
		})
		dsl.Payload(func() {
			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})
			dsl.Field(1, "id", dsl.Int64, func() {
				dsl.Description("the user id to be updated")
			})
			dsl.Field(2, "user", UserPayload, "")

			dsl.Required("id", "user", "token")
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
		dsl.Description("delete existing user")
		dsl.Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			dsl.Scope("delete-user") // Enforce presence of "list-users" scope in JWT claims.
		})
		dsl.Payload(func() {
			dsl.TokenField(1, "token", dsl.String, func() {
				dsl.Description("JWT used for authentication")
			})
			dsl.Field(2, "id", dsl.Int64, func() {
				dsl.Description("the user id to be updated")
			})
			dsl.Required("token", "id")
		})

		dsl.HTTP(func() {
			dsl.DELETE("/{id}")
			dsl.Response(dsl.StatusOK)
			dsl.Response("not-exist", dsl.StatusNotFound)
		})
	})
	dsl.Method("login", func() {
		dsl.Description("delete existing user")

		dsl.Payload(func() {

			dsl.Field(1, "UserName", dsl.String, func() {
				dsl.Description("user name")
			})
			dsl.Field(2, "USerPassword", dsl.String, func() {
				dsl.Description("user password")
			})
			dsl.Required("UserName", "USerPassword")
		})
		dsl.Result(LoginRes)
		dsl.HTTP(func() {
			dsl.POST("/login")
			dsl.Response(dsl.StatusOK)
			dsl.Response("not-exist", dsl.StatusNotFound)
		})
	})

})

// ListUserResp ...
var ListUserResp = dsl.Type("ListUserResp", func() {
	dsl.Field(1, "UserID", dsl.String, "user Id")
	dsl.Field(2, "UserName", dsl.String, "user name")
	dsl.Field(4, "UserDeposit", dsl.Int, "user name")
	dsl.Field(5, "UserRole", dsl.String, "user role")
	dsl.Required("UserID", "UserName", "UserDeposit", "UserRole")

})

// UserPayload ...
var UserPayload = dsl.Type("UserPayload", func() {
	dsl.Field(1, "UserName", dsl.String, "user name")
	dsl.Field(2, "UserPassword", dsl.String, "user passowrd")
	dsl.Field(4, "UserRole", dsl.String, func() {
		dsl.Enum(models.Seller.String(), models.Buyer.String())
	})
	dsl.Required("UserName", "UserPassword", "UserRole")
})

// LoginRes ...
var LoginRes = dsl.Type("LoginRes", func() {
	dsl.Field(1, "token", dsl.String, "user token")
	dsl.Field(2, "id", dsl.Int, "user ID")
	dsl.Required("token", "id")
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
