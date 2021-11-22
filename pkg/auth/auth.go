package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// AuthI ...
type AuthI interface {
	CreateToken() (AuthI, error)
	Interval(interval int) AuthI
	UserID(userID int64) AuthI
	Role(role string) AuthI

	VerifyToken(tokenString string) (AuthI, error)
	IsAdmin() bool
	IsUser() bool
	GetUserID() int64
	GetToken() string
}

// AuthBuilder ...
type AuthBuilder struct {
	interval    int
	accessToken string
	jwtAccess   string
	role        string
	expire      int
	userID      int64
	valid       bool
}

// NewAuth ...
func NewAuth() (AuthI, error) {
	jwtAccess := os.Getenv("JWT_ACCESS_SECRET")
	if jwtAccess == "" {
		return nil, fmt.Errorf("Invalid env variable")
	}
	return &AuthBuilder{jwtAccess: jwtAccess}, nil
}

// Interval ...
func (a *AuthBuilder) Interval(interval int) AuthI {
	a.interval = interval
	return a
}

// UserID ...
func (a *AuthBuilder) UserID(userID int64) AuthI {
	a.userID = userID
	return a
}

// Role ...
func (a *AuthBuilder) Role(role string) AuthI {
	a.role = role
	return a
}

// CreateToken ...
func (a *AuthBuilder) CreateToken() (AuthI, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["role"] = a.role
	atClaims["user_id"] = a.userID
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(a.interval)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	a.accessToken, err = at.SignedString([]byte(a.jwtAccess))
	if err != nil {
		return a, err
	}
	return a, nil
}

// VerifyToken ...
func (a *AuthBuilder) VerifyToken(tokenString string) (AuthI, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.jwtAccess), nil
	})
	if err != nil {
		return a, err
	}

	a.valid = token.Valid
	claims, _ := token.Claims.(jwt.MapClaims)

	a.userID = int64(claims["user_id"].(float64))
	a.role = claims["role"].(string)
	return a, nil
}

// IsAdmin ...
func (a *AuthBuilder) IsAdmin() bool {
	return a.role == string(Admin)
}

// IsUser ...
func (a *AuthBuilder) IsUser() bool {
	return a.role == string(User)
}

// GetUser ...
func (a *AuthBuilder) GetUserID() int64 {
	return a.userID
}

// GetAccessToken ...
func (a *AuthBuilder) GetToken() string {
	return a.accessToken
}
