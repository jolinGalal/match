package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Builder ...
type Builder struct {
	host     string
	password string
	port     int
	dbName   string
	username string
	create   bool
	sslmode  string
	migrate  bool
	models   []interface{}
}

// New ...
func New() *Builder {
	return &Builder{}
}

// Host ...
func (b *Builder) Host(host string) *Builder {
	b.host = host
	return b
}

// Sslmode ...
func (b *Builder) Sslmode(sslmode string) *Builder {
	b.sslmode = sslmode
	return b
}

// DbName ...
func (b *Builder) DbName(name string) *Builder {
	b.dbName = name
	return b
}

// UserName ...
func (b *Builder) UserName(name string) *Builder {
	b.username = name
	return b
}

// Password ...
func (b *Builder) Password(password string) *Builder {
	b.password = password
	return b
}

// Port ...
func (b *Builder) Port(port int) *Builder {
	b.port = port
	return b
}

// Models ...
func (b *Builder) Models(model interface{}) *Builder {
	b.models = append(b.models, model)
	return b
}

// Migrate ...
func (b *Builder) Migrate(migrate bool) *Builder {
	b.migrate = migrate
	return b
}

// Create ...
func (b *Builder) Create(create bool) *Builder {
	b.create = create
	return b
}

// Get builds the database instance
func (b *Builder) Get() (*gorm.DB, error) {

	if len(b.host) == 0 {
		return nil, fmt.Errorf("Invalid host name")
	}

	if len(b.dbName) == 0 {
		return nil, fmt.Errorf("Invalid database name")
	}

	if len(b.username) == 0 {
		return nil, fmt.Errorf("Invalid username")
	}

	if len(b.password) == 0 {
		return nil, fmt.Errorf("Invalid password")
	}

	if b.create {
		err := b.CreateDB()
		if err != nil {
			return nil, err
		}
	}
	db, err := b.Migration()
	if err != nil {
		return nil, err
	}
	var query = `ALTER TABLE Products DROP CONSTRAINT IF EXISTS constraint_user;`
	err = db.Exec(query).Error
	if err != nil {
		return nil, err
	}
	query = `ALTER TABLE Products ADD CONSTRAINT constraint_user FOREIGN KEY (seller_id) REFERENCES users(user_id) ON DELETE SET NULL ON UPDATE CASCADE;`
	err = db.Exec(query).Error
	return db, err
}

// CreateDB ...
func (b *Builder) CreateDB() (err error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d   sslmode=%s", b.username, b.password, b.host, b.port, b.sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", b.dbName)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		return rs.Error
	}

	// if not create it
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", b.dbName)
		err = db.Exec(stmt).Error
	}
	return err
}

//conn connect to database
func (b *Builder) conn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d  dbname =%s sslmode=%s", b.username, b.password, b.host, b.port, b.dbName, b.sslmode)
	dB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dB, nil
}

// Migrate models to the dtaabase...
func (b *Builder) Migration() (*gorm.DB, error) {
	db, err := b.conn()
	if err != nil {
		return nil, err
	}
	if b.migrate {
		for _, table := range b.models {
			db.Migrator()
			if !db.Migrator().HasTable(table) {
				db.Migrator().CreateTable(table)
			}

			db.AutoMigrate(table)
		}
	}

	return db, err
}
