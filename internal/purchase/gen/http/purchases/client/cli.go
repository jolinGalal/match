// Code generated by goa v3.5.2, DO NOT EDIT.
//
// purchases HTTP client CLI support package
//
// Command:
// $ goa gen github.com/jolinGalal/match/internal/purchase/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	purchases "github.com/jolinGalal/match/internal/purchase/gen/purchases"
	goa "goa.design/goa/v3/pkg"
)

// BuildDepositPayload builds the payload for the purchases deposit endpoint
// from CLI flags.
func BuildDepositPayload(purchasesDepositBody string, purchasesDepositID string, purchasesDepositToken string) (*purchases.DepositPayload, error) {
	var err error
	var body DepositRequestBody
	{
		err = json.Unmarshal([]byte(purchasesDepositBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"deposit\": 5\n   }'")
		}
		if !(body.Deposit == 5 || body.Deposit == 10 || body.Deposit == 20 || body.Deposit == 50 || body.Deposit == 100) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.deposit", body.Deposit, []interface{}{5, 10, 20, 50, 100}))
		}
		if err != nil {
			return nil, err
		}
	}
	var id int64
	{
		id, err = strconv.ParseInt(purchasesDepositID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var token string
	{
		token = purchasesDepositToken
	}
	v := &purchases.DepositPayload{
		Deposit: body.Deposit,
	}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildBuyPayload builds the payload for the purchases buy endpoint from CLI
// flags.
func BuildBuyPayload(purchasesBuyBody string, purchasesBuyID string, purchasesBuyToken string) (*purchases.BuyPayload, error) {
	var err error
	var body BuyRequestBody
	{
		err = json.Unmarshal([]byte(purchasesBuyBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"purchase\": [\n         {\n            \"productID\": 7749138134878224859,\n            \"purductAmount\": 6051214521179394048\n         },\n         {\n            \"productID\": 7749138134878224859,\n            \"purductAmount\": 6051214521179394048\n         },\n         {\n            \"productID\": 7749138134878224859,\n            \"purductAmount\": 6051214521179394048\n         },\n         {\n            \"productID\": 7749138134878224859,\n            \"purductAmount\": 6051214521179394048\n         }\n      ]\n   }'")
		}
		if body.Purchase == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("purchase", "body"))
		}
		for _, e := range body.Purchase {
			if e != nil {
				if err2 := ValidateProductsListRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id int64
	{
		id, err = strconv.ParseInt(purchasesBuyID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var token string
	{
		token = purchasesBuyToken
	}
	v := &purchases.BuyPayload{}
	if body.Purchase != nil {
		v.Purchase = make([]*purchases.ProductsList, len(body.Purchase))
		for i, val := range body.Purchase {
			v.Purchase[i] = marshalProductsListRequestBodyToPurchasesProductsList(val)
		}
	}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildResetPayload builds the payload for the purchases reset endpoint from
// CLI flags.
func BuildResetPayload(purchasesResetID string, purchasesResetToken string) (*purchases.ResetPayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(purchasesResetID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var token string
	{
		token = purchasesResetToken
	}
	v := &purchases.ResetPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}
