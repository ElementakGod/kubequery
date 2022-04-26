package database

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	gsql "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type SqliteDatabase struct {
	DBName string
}

func NewSqliteDatabase(dbName string) Database {
	return &SqliteDatabase{dbName}
}

func (s *SqliteDatabase) Open(opts *DbOptions) (db *gorm.DB, err error) {
	if opts == nil {
		err = fmt.Errorf("database options cannot be empty")
		return
	}

	db, err = gorm.Open(
		gsql.Open(opts.DbConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}

func (s *SqliteDatabase) GetConnection() string {
	return s.DBName
}
