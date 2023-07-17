package db

import (
	"github.com/knadh/koanf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB(conf *koanf.Koanf) *gorm.DB {

	dsn := conf.String("database.host")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	return db

}
