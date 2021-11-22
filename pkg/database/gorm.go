package database

import (
	"strings"

	"gorm.io/gorm"
)

//GormI ...
type GormI interface {
	Rest() GormI
	Select(str ...string) GormI
	SelectDistinct(str ...string) GormI
	LeftJoin(tableName, fOperator, sOperator string) GormI
	RightJoin(tableName, fOperator, sOperator string) GormI
	Join(tableName, fOperator, sOperator string) GormI
	Where(condition string) GormI
	Model(model interface{}) GormI
	Paging(pageNumber, pageSize *int) GormI
	Sort(str ...string) GormI
	Count() (*int64, error)
	Find(res interface{}) error
	Save(i interface{}) error
	GetDB() *gorm.DB
	Delete(table, id interface{}) error
	Execute(str string) error
}

//GormBuilder ...
type GormBuilder struct {
	DB   *gorm.DB
	conn *gorm.DB
}

//NewDB ...
func NewDB(db *gorm.DB) GormI {
	return &GormBuilder{DB: db, conn: db}

}

// GetDB ...
func (g *GormBuilder) GetDB() *gorm.DB {
	return g.DB
}

// Select ...
func (g *GormBuilder) Select(str ...string) GormI {
	g.DB = g.DB.Select(strings.Join(str, ","))
	return g
}

// SelectDistinct ...
func (g *GormBuilder) SelectDistinct(str ...string) GormI {
	g.DB = g.DB.Select("DISTINCT " + strings.Join(str, ","))
	return g
}

// LeftJoin ...
func (g *GormBuilder) LeftJoin(tableName, fOperator, sOperator string) GormI {
	g.DB = g.DB.Joins("left join " + tableName + " on " + fOperator + " = " + sOperator)
	return g
}

// RightJoin ...
func (g *GormBuilder) RightJoin(tableName, fOperator, sOperator string) GormI {
	g.DB = g.DB.Joins("right join " + tableName + " on " + fOperator + " = " + sOperator)
	return g
}

// Join ...
func (g *GormBuilder) Join(tableName, fOperator, sOperator string) GormI {
	g.DB = g.DB.Joins("join " + tableName + " on " + fOperator + " = " + sOperator)
	return g
}

// Where ...
func (g *GormBuilder) Where(condition string) GormI {
	g.DB = g.DB.Where(condition)
	return g
}

// Model ...
func (g *GormBuilder) Model(model interface{}) GormI {
	g.DB = g.DB.Model(model)
	return g
}

// Paging ...
func (g *GormBuilder) Paging(pageNumber, pageSize *int) GormI {
	if pageSize != nil && pageNumber != nil {
		if *pageSize != -1 && *pageNumber != -1 {
			offset := (*pageNumber - 1) * *pageSize
			g.DB = g.DB.Limit(*pageSize).Offset(offset)
		}
	}
	return g
}

// Sort ...
func (g *GormBuilder) Sort(str ...string) GormI {
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
	return g.DB.Save(i).Error
}

// Detele ...
func (g *GormBuilder) Delete(table, id interface{}) error {
	return g.DB.Delete(table, id).Error
}

// Execute ...
func (g *GormBuilder) Execute(str string) error {
	return g.DB.Exec(str).Error
}

// Rest ...
func (g *GormBuilder) Rest() GormI {
	g.DB = g.conn
	return g
}
