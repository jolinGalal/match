package user

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jolinGalal/match/pkg/auth"
	"github.com/stretchr/testify/assert"
)

// obj of customer service
var obj userssrvc

//newUserService ...
func newUserService(auth auth.AuthI) {
	logger := log.New(os.Stderr, "[match] ", log.Ltime)
	obj.db = NewDBFunc()
	obj.auth = auth
	obj.logger = logger
}

//teardown run after all tests get executed
func teardown() {
}

//TestJWTAuth ...
func TestJWTAuth(t *testing.T) {

	VerifyTokenFunc = func(tokenString string) (auth.AuthI, error) {

		return nil, nil
	}
	auth := &authMock{}
	newUserService(auth)
	_, err := obj.JWTAuth(context.Background(), "mock", nil)
	assert.Nil(t, err)
}

//TestJWTAuthFail ...
func TestJWTAuthFail(t *testing.T) {
	VerifyTokenFunc = func(tokenString string) (auth.AuthI, error) {

		return nil, fmt.Errorf("error")
	}

	auth := &authMock{}
	newUserService(auth)
	_, err := obj.JWTAuth(context.Background(), "mock", nil)
	assert.NotNil(t, err)
}

//TestJWTAuthFailEmptyStr ...
func TestJWTAuthFailEmptyStr(t *testing.T) {
	VerifyTokenFunc = func(tokenString string) (auth.AuthI, error) {

		return nil, fmt.Errorf("error")
	}

	auth := &authMock{}
	newUserService(auth)
	_, err := obj.JWTAuth(context.Background(), "", nil)
	assert.NotNil(t, err)
}
