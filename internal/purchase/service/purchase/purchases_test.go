package purchase

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jolinGalal/match/pkg/auth"
	"github.com/jolinGalal/match/pkg/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// obj of customer service
var obj purchasessrvc

var authGetUser = `SELECT count(*) FROM "users" WHERE users.user_id = 1 AND "users"."deleted_at" IS NULL`

//newPurchaseService ...
func newPurchaseService(auth auth.AuthI) {
	logger := log.New(os.Stderr, "[match] ", log.Ltime)
	obj.db = NewDBFunc()
	obj.auth = auth
	obj.logger = logger
}

//teardown run after all tests get executed
func teardown() {
}

// JWTNewDB ...
func JWTNewDB() {
	NewDBFunc = func() database.GormI {
		db, mock, err := sqlmock.New()
		if err != nil {
			log.Fatalln(err)
		}

		g, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
		mock.ExpectQuery(regexp.QuoteMeta(authGetUser)).WillReturnRows(rows)
		return &GormBuilder{DB: g, Mock: mock}

	}
}

//TestJWTAuth ...
func TestJWTAuth(t *testing.T) {
	authObj := &authMock{}

	VerifyTokenFunc = func(tokenString string) (auth.AuthI, error) {

		return authObj, nil
	}
	IsUserFunc = func() bool {
		return true
	}
	GetUserIDFunc = func() int64 {
		return 1
	}
	JWTNewDB()

	newPurchaseService(authObj)
	_, err := obj.JWTAuth(context.Background(), "mock", nil)
	assert.Nil(t, err)
}

//TestJWTAuthFail ...
func TestJWTAuthFail(t *testing.T) {
	authObj := &authMock{}

	VerifyTokenFunc = func(tokenString string) (auth.AuthI, error) {

		return nil, fmt.Errorf("error")
	}
	IsUserFunc = func() bool {
		return true
	}
	GetUserIDFunc = func() int64 {
		return 1
	}
	JWTNewDB()
	newPurchaseService(authObj)
	_, err := obj.JWTAuth(context.Background(), "mock", nil)
	assert.NotNil(t, err)
}

//TestJWTAuthFailEmptyStr ...
func TestJWTAuthFailEmptyStr(t *testing.T) {
	authObj := &authMock{}
	VerifyTokenFunc = func(tokenString string) (auth.AuthI, error) {

		return nil, fmt.Errorf("error")
	}
	IsUserFunc = func() bool {
		return true
	}
	GetUserIDFunc = func() int64 {
		return 1
	}
	JWTNewDB()

	newPurchaseService(authObj)
	_, err := obj.JWTAuth(context.Background(), "", nil)
	assert.NotNil(t, err)
}
