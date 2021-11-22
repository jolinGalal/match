package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestConvertStringToRoleSuccess ...
func TestConvertStringToRoleSuccess(t *testing.T) {
	user := NewUser()
	r, err := user.ConvertStringToRole("Seller")

	assert.Nil(t, err)
	assert.Equal(t, r, Seller)
}

//TestConvertStringToRoleFail ...
func TestConvertStringToRoleFail(t *testing.T) {
	user := NewUser()
	r, err := user.ConvertStringToRole("not role")

	assert.NotNil(t, err)
	assert.Equal(t, r, "")
}

//TestValidDepositeSuccess ...
func TestValidDepositeSuccess(t *testing.T) {
	user := NewUser()
	v := user.ValidDeposit(10)

	assert.Equal(t, v, true)
}

//TestValidDepositeFail ...
func TestValidDepositeFail(t *testing.T) {
	user := NewUser()
	v := user.ValidDeposit(16)

	assert.Equal(t, v, false)
}
