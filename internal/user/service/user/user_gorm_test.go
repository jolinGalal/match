package user

import (
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jolinGalal/match/pkg/database"
	"gorm.io/gorm"
)

//GormBuilder ...
type GormBuilder struct {
	Mock     sqlmock.Sqlmock
	MockConn sqlmock.Sqlmock
	DB       *gorm.DB
}

var NewDBFunc func() database.GormI

// GetDB ...
func (g *GormBuilder) GetDB() *gorm.DB {
	return g.DB
}

// Select ...
func (g *GormBuilder) Select(str ...string) database.GormI {
	g.DB = g.DB.Select(strings.Join(str, ","))
	return g
}

// SelectDistinct ...
func (g *GormBuilder) SelectDistinct(str ...string) database.GormI {
	g.DB = g.DB.Select("DISTINCT " + strings.Join(str, ","))
	return g
}

// LeftJoin ...
func (g *GormBuilder) LeftJoin(tableName, fOperator, sOperator string) database.GormI {
	g.DB = g.DB.Joins("left join " + tableName + " on " + fOperator + " = " + sOperator)
	return g
}

// RightJoin ...
func (g *GormBuilder) RightJoin(tableName, fOperator, sOperator string) database.GormI {
	g.DB = g.DB.Joins("right join " + tableName + " on " + fOperator + " = " + sOperator)
	return g
}

// Join ...
func (g *GormBuilder) Join(tableName, fOperator, sOperator string) database.GormI {
	g.DB = g.DB.Joins("join " + tableName + " on " + fOperator + " = " + sOperator)
	return g
}

// Where ...
func (g *GormBuilder) Where(condition string) database.GormI {
	g.DB = g.DB.Where(condition)
	return g
}

// Model ...
func (g *GormBuilder) Model(model interface{}) database.GormI {
	g.DB = g.DB.Model(model)
	return g
}

// Paging ...
func (g *GormBuilder) Paging(pageNumber, pageSize *int) database.GormI {
	if pageSize != nil && pageNumber != nil {
		if *pageSize != -1 && *pageNumber != -1 {
			offset := (*pageNumber - 1) * *pageSize
			g.DB = g.DB.Limit(*pageSize).Offset(offset)
		}
	}
	return g
}

// Sort ...
func (g *GormBuilder) Sort(str ...string) database.GormI {
	g.DB = g.DB.Order(strings.Join(str, ","))
	return g
}

// Count ...
func (g *GormBuilder) Count() (*int64, error) {
	totalCount := new(int64)
	err := g.DB.Count(totalCount).Error
	return totalCount, err
}

// Find ...
func (g *GormBuilder) Find(res interface{}) error {
	return g.DB.Find(res).Error
}

// Save ...
func (g *GormBuilder) Save(i interface{}) error {
	return g.DB.Create(i).Error
}

// Detele ...
func (g *GormBuilder) Delete(table, id interface{}) error {
	return g.DB.Delete(table, id).Error
}

// Rest ...
func (g *GormBuilder) Rest() database.GormI {
	return g
}
