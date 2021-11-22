// Code generated by goa v3.5.2, DO NOT EDIT.
//
// purchases client HTTP transport
//
// Command:
// $ goa gen github.com/jolinGalal/match/internal/purchase/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the purchases service endpoint HTTP clients.
type Client struct {
	// Deposit Doer is the HTTP client used to make requests to the deposit
	// endpoint.
	DepositDoer goahttp.Doer

	// Buy Doer is the HTTP client used to make requests to the buy endpoint.
	BuyDoer goahttp.Doer

	// Reset Doer is the HTTP client used to make requests to the reset endpoint.
	ResetDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the purchases service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		DepositDoer:         doer,
		BuyDoer:             doer,
		ResetDoer:           doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Deposit returns an endpoint that makes HTTP requests to the purchases
// service deposit server.
func (c *Client) Deposit() goa.Endpoint {
	var (
		encodeRequest  = EncodeDepositRequest(c.encoder)
		decodeResponse = DecodeDepositResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDepositRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DepositDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("purchases", "deposit", err)
		}
		return decodeResponse(resp)
	}
}

// Buy returns an endpoint that makes HTTP requests to the purchases service
// buy server.
func (c *Client) Buy() goa.Endpoint {
	var (
		encodeRequest  = EncodeBuyRequest(c.encoder)
		decodeResponse = DecodeBuyResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildBuyRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.BuyDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("purchases", "buy", err)
		}
		return decodeResponse(resp)
	}
}

// Reset returns an endpoint that makes HTTP requests to the purchases service
// reset server.
func (c *Client) Reset() goa.Endpoint {
	var (
		encodeRequest  = EncodeResetRequest(c.encoder)
		decodeResponse = DecodeResetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildResetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ResetDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("purchases", "reset", err)
		}
		return decodeResponse(resp)
	}
}
