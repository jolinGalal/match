// Code generated by goa v3.5.2, DO NOT EDIT.
//
// users HTTP client CLI support package
//
// Command:
// $ goa gen github.com/jolinGalal/match/internal/user/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	users "github.com/jolinGalal/match/internal/user/gen/users"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the users list endpoint from CLI
// flags.
func BuildListPayload(usersListSortDirection string, usersListSortKey string, usersListPageNumber string, usersListPageSize string, usersListToken string) (*users.ListPayload, error) {
	var err error
	var sortDirection string
	{
		if usersListSortDirection != "" {
			sortDirection = usersListSortDirection
			if !(sortDirection == "asc" || sortDirection == "desc") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("sortDirection", sortDirection, []interface{}{"asc", "desc"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var sortKey string
	{
		if usersListSortKey != "" {
			sortKey = usersListSortKey
			if !(sortKey == "CreatedAt" || sortKey == "UserID" || sortKey == "UserName") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("sortKey", sortKey, []interface{}{"CreatedAt", "UserID", "UserName"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var pageNumber int
	{
		if usersListPageNumber != "" {
			var v int64
			v, err = strconv.ParseInt(usersListPageNumber, 10, 64)
			pageNumber = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for pageNumber, must be INT")
			}
		}
	}
	var pageSize int
	{
		if usersListPageSize != "" {
			var v int64
			v, err = strconv.ParseInt(usersListPageSize, 10, 64)
			pageSize = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for pageSize, must be INT")
			}
		}
	}
	var token string
	{
		token = usersListToken
	}
	v := &users.ListPayload{}
	v.SortDirection = sortDirection
	v.SortKey = sortKey
	v.PageNumber = pageNumber
	v.PageSize = pageSize
	v.Token = token

	return v, nil
}

// BuildCreatePayload builds the payload for the users create endpoint from CLI
// flags.
func BuildCreatePayload(usersCreateBody string) (*users.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(usersCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"user\": {\n         \"UserName\": \"Est sed in.\",\n         \"UserPassword\": \"Voluptates ut provident in eos sit quo.\",\n         \"UserRole\": \"seller\"\n      }\n   }'")
		}
		if body.User == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
		}
		if body.User != nil {
			if err2 := ValidateUserPayloadRequestBody(body.User); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	v := &users.CreatePayload{}
	if body.User != nil {
		v.User = marshalUserPayloadRequestBodyToUsersUserPayload(body.User)
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the users update endpoint from CLI
// flags.
func BuildUpdatePayload(usersUpdateBody string, usersUpdateID string, usersUpdateToken string) (*users.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(usersUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"user\": {\n         \"UserName\": \"Est sed in.\",\n         \"UserPassword\": \"Voluptates ut provident in eos sit quo.\",\n         \"UserRole\": \"seller\"\n      }\n   }'")
		}
		if body.User == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
		}
		if body.User != nil {
			if err2 := ValidateUserPayloadRequestBody(body.User); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id int64
	{
		id, err = strconv.ParseInt(usersUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var token string
	{
		token = usersUpdateToken
	}
	v := &users.UpdatePayload{}
	if body.User != nil {
		v.User = marshalUserPayloadRequestBodyToUsersUserPayload(body.User)
	}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildDeletePayload builds the payload for the users delete endpoint from CLI
// flags.
func BuildDeletePayload(usersDeleteID string, usersDeleteToken string) (*users.DeletePayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(usersDeleteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var token string
	{
		token = usersDeleteToken
	}
	v := &users.DeletePayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildLoginPayload builds the payload for the users login endpoint from CLI
// flags.
func BuildLoginPayload(usersLoginBody string) (*users.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(usersLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"USerPassword\": \"Expedita qui velit eveniet atque.\",\n      \"UserName\": \"Quia rem debitis corporis esse.\"\n   }'")
		}
	}
	v := &users.LoginPayload{
		UserName:     body.UserName,
		USerPassword: body.USerPassword,
	}

	return v, nil
}
