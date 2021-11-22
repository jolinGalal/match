package purchase

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	purchases "github.com/jolinGalal/match/internal/purchase/gen/purchases"
	"github.com/jolinGalal/match/pkg/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var deposit = `SELECT * FROM "users" WHERE  users.user_id = 1 AND "users"."deleted_at" IS NULL`
var depositUpdate = `update users set user_deposit=15 where users.user_id=1`

// TestDeposite ...
func TestDeposite(t *testing.T) {
	var authObj = newAuth()
	p := &purchases.DepositPayload{
		Token:   "token",
		ID:      1,
		Deposit: 5,
	}
	GetUserIDFunc = func() int64 {
		return 1
	}
	NewDBFunc = func() database.GormI {

		db, mock, err := sqlmock.New()
		if err != nil {
			log.Fatalln(err)
		}

		g, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		rows := sqlmock.NewRows([]string{"user_id", "user_name", "user_password", "user_deposit", "user_role"}).
			AddRow(1, "name", "pass", 10, "seller")
		mock.ExpectQuery(regexp.QuoteMeta(deposit)).WillReturnRows(rows)

		mock.ExpectExec(regexp.QuoteMeta(depositUpdate)).WillReturnResult(sqlmock.NewResult(0, 1))
		// .
		// WithArgs(sqlmock.AnyArg(), "purchase", "_VLg7HIChNQ9RIEOCkrprWD450M=").
		// WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		return &GormBuilder{DB: g, Mock: mock}

	}

	newPurchaseService(authObj)
	err := obj.Deposit(context.Background(), p)
	assert.Nil(t, err)

}

// TestDepositeFail ...
func TestDepositeFail(t *testing.T) {
	var authObj = newAuth()
	p := &purchases.DepositPayload{
		Token:   "token",
		ID:      1,
		Deposit: 5,
	}
	GetUserIDFunc = func() int64 {
		return 1
	}
	NewDBFunc = func() database.GormI {

		db, mock, err := sqlmock.New()
		if err != nil {
			log.Fatalln(err)
		}

		g, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		mock.ExpectQuery(regexp.QuoteMeta(deposit)).WillReturnError(fmt.Errorf("error"))

		mock.ExpectExec(regexp.QuoteMeta(depositUpdate)).WillReturnResult(sqlmock.NewResult(0, 1))

		return &GormBuilder{DB: g, Mock: mock}

	}

	newPurchaseService(authObj)
	err := obj.Deposit(context.Background(), p)
	assert.NotNil(t, err)

}
