package user

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	users "github.com/jolinGalal/match/internal/user/gen/users"
	"github.com/jolinGalal/match/pkg/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var userCount = `SELECT count(*) FROM "users" WHERE lower(trim(users.user_name)) like lower(trim('user')) AND "users"."deleted_at" IS NULL`
var userInsert = `INSERT INTO "users" (user_name,user_password,userrole) VALUES ('user','Password@122','seller')`

// TestCreate ...
func TestCreate(t *testing.T) {
	var authObj = newAuth()
	p := &users.CreatePayload{
		User: &users.UserPayload{
			UserName:     "user",
			UserPassword: "Password@122",
			UserRole:     "seller",
		},
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
		rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
		mock.ExpectQuery(regexp.QuoteMeta(userCount)).WillReturnRows(rows)
		//mock.ExpectExec(userInsert)
		mock.ExpectQuery(`INSERT INTO "table" (.+) RETURNING`).
			WithArgs(sqlmock.AnyArg(), "user", "_VLg7HIChNQ9RIEOCkrprWD450M=").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		return &GormBuilder{DB: g, Mock: mock}

	}

	newUserService(authObj)
	err := obj.Create(context.Background(), p)
	assert.Nil(t, err)

}

// TestCreateFail ...
func TestCreateFail(t *testing.T) {
	var authObj = newAuth()
	p := &users.CreatePayload{
		User: &users.UserPayload{
			UserName:     "user",
			UserPassword: "Password@122",
			UserRole:     "seller",
		},
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
		rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
		mock.ExpectQuery(regexp.QuoteMeta(userCount)).WillReturnRows(rows).WillReturnError(fmt.Errorf("Error"))

		return &GormBuilder{DB: g, Mock: mock}

	}

	newUserService(authObj)
	err := obj.Create(context.Background(), p)
	assert.NotNil(t, err)

}
