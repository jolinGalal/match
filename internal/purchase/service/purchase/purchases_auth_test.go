package purchase

import (
	"github.com/jolinGalal/match/pkg/auth"
)

//verifyTokenFunc ....
var CreateTokenFunc func() (auth.AuthI, error)
var IntervalFunc func(interval int) auth.AuthI
var UserIDFunc func(userID int64) auth.AuthI
var RoleFunc func(role string) auth.AuthI
var VerifyTokenFunc func(tokenString string) (auth.AuthI, error)
var IsAdminFunc func() bool
var IsUserFunc func() bool
var GetUserIDFunc func() int64
var GetTokenFunc func() string

//authMock ...
type authMock struct {
}

func newAuth() auth.AuthI {
	return &authMock{}
}

//VerifyToken ...
func (a *authMock) VerifyToken(tokenString string) (auth.AuthI, error) {

	return VerifyTokenFunc(tokenString)
}

func (a *authMock) CreateToken() (auth.AuthI, error) {
	return CreateTokenFunc()
}
func (a *authMock) Interval(interval int) auth.AuthI {
	return IntervalFunc(interval)
}

func (a *authMock) UserID(userID int64) auth.AuthI {
	return UserIDFunc(userID)
}
func (a *authMock) Role(role string) auth.AuthI {
	return RoleFunc(role)
}

func (a *authMock) IsAdmin() bool {
	return IsAdminFunc()
}
func (a *authMock) IsUser() bool {
	return IsUserFunc()
}
func (a *authMock) GetUserID() int64 {
	return GetUserIDFunc()
}
func (a *authMock) GetToken() string {
	return GetTokenFunc()
}
