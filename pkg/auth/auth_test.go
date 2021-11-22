package auth

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//TestAuthSuccess ...
func TestAuthSuccess(t *testing.T) {
	os.Setenv("JWT_ACCESS_SECRET", "jwttokenkeys")
	auth, err := NewAuth()
	assert.Nil(t, err)
	auth = auth.Interval(3).Role("seller").UserID(1)
	auth, err = auth.CreateToken()
	assert.Nil(t, err)
	auth, err = auth.VerifyToken(auth.GetToken())
	assert.Nil(t, err)
}

//TestAuthFail ...
func TestAuthFail(t *testing.T) {
	os.Setenv("JWT_ACCESS_SECRET", "jwttokenkeys")
	auth, err := NewAuth()
	assert.Nil(t, err)
	auth = auth.Interval(1).Role("seller").UserID(1)
	auth, err = auth.CreateToken()
	assert.Nil(t, err)
	time.Sleep(2 * time.Minute)
	auth, err = auth.VerifyToken(auth.GetToken())
	assert.NotNil(t, err)
}
