package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"moul.io/zapgorm2"
)
import "gorm.io/driver/postgres"

var _db *gorm.DB

type Scopeb func(*gorm.DB) *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		"127.0.0.1",
		"directus",
		"directus",
		"directus",
		"5432",
	)
	zl, _ := zap.NewDevelopment()
	l := zapgorm2.New(zl)
	l.SetAsDefault()
	ll := l.LogMode(logger.Info)
	log.Default().Println("[InitDatabase] connecting to database with dsn:" + dsn)
	var err error
	_db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: ll,
	})
	if err != nil {
		log.Fatalf("connection to database failed: %v", err)
	}

}

func DB() *gorm.DB {
	if _db == nil {
		InitDatabase()
	}
	db := _db.Debug()
	return db
}

//func Create[T any](data *T) {
//	err := DB().Create(&data).Error
//	if err != nil {
//		panic(err)
//	}
//}
