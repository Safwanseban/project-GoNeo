package utils

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB(db *sql.DB) *gorm.DB {

	gormHandler, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		return nil
	}

	return gormHandler
}
