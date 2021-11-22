// Code generated by goa v3.5.2, DO NOT EDIT.
//
// users client
//
// Command:
// $ goa gen github.com/jolinGalal/match/internal/user/design

package users

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "users" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	CreateEndpoint goa.Endpoint
	UpdateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
	LoginEndpoint  goa.Endpoint
}

// NewClient initializes a "users" service client given the endpoints.
func NewClient(list, create, update, delete_, login goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		CreateEndpoint: create,
		UpdateEndpoint: update,
		DeleteEndpoint: delete_,
		LoginEndpoint:  login,
	}
}

// List calls the "list" endpoint of the "users" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res *ListResult, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListResult), nil
}

// Create calls the "create" endpoint of the "users" service.
func (c *Client) Create(ctx context.Context, p *CreatePayload) (err error) {
	_, err = c.CreateEndpoint(ctx, p)
	return
}

// Update calls the "update" endpoint of the "users" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}

// Delete calls the "delete" endpoint of the "users" service.
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}

// Login calls the "login" endpoint of the "users" service.
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *LoginRes, err error) {
	var ires interface{}
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginRes), nil
}