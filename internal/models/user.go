package models

import (
	"fmt"
	"strings"
	"sync"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	UserID       int64 `gorm:"primaryKey;unique;autoIncrement:true"`
	UserName     string
	UserPassword string
	UserDeposit  int
	UserRole     Role
}

//UserI ...
type UserI interface {
	GetTableName() string
	ConvertStringToRole(string) (Role, error)
	ValidDeposit(int) bool
	ID() string
	Name() string
	Password() string
	Deposite() string
	Role() string
}

// UserLock ...
var UserLock = &sync.Mutex{}

// userInstance ...
var userInstance *User

// NewUser ...
func NewUser() UserI {
	if userInstance == nil {
		UserLock.Lock()
		defer UserLock.Unlock()
		if userInstance == nil {
			userInstance = &User{}
		}
	}

	return userInstance
}

// Role
type Role string

const (
	// seller role ...
	Seller Role = "seller"
	// buyer role ...
	Buyer Role = "buyer"
)

//String convert role enum to string
func (r Role) String() string {
	switch r {
	case Seller:
		return "seller"
	case Buyer:
		return "buyer"
	default:
		return ""
	}
}

//Deposit value ...
var Deposit = map[int]bool{
	100: true,
	50:  true,
	20:  true,
	10:  true,
	5:   true,
}

// GetTableName ...
func (u *User) GetTableName() string {
	return "users"
}

// ConvertStringToRole ...
func (u *User) ConvertStringToRole(roleStr string) (Role, error) {
	switch strings.ToLower(roleStr) {
	case Seller.String():
		return Seller, nil
	case Buyer.String():
		return Buyer, nil
	default:
		return "", fmt.Errorf("Invalid Role")
	}
}

// ValidDeposit ...
func (u *User) ValidDeposit(value int) bool {
	exits, _ := Deposit[value]
	return exits
}

// ID ...
func (u *User) ID() string {
	return fmt.Sprintf("%s.user_id", u.GetTableName())
}

// Name ...
func (u *User) Name() string {
	return fmt.Sprintf("%s.user_name", u.GetTableName())
}

// Password ...
func (u *User) Password() string {
	return fmt.Sprintf("%s.user_password", u.GetTableName())
}

//  Deposite...
func (u *User) Deposite() string {
	return fmt.Sprintf("%s.user_deposit", u.GetTableName())
}

//  Role...
func (u *User) Role() string {
	return fmt.Sprintf("%s.user_role", u.GetTableName())
}

//  CreatedAt...
func (u *User) CreatedAt() string {
	return fmt.Sprintf("%s.created_at", u.GetTableName())
}
