package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbOptions struct {
	DbType   string
	DbConn   string
	LogLevel logger.LogLevel
}

type Database interface {
	Open(opts *DbOptions) (db *gorm.DB, err error)
	GetConnection() string
}
