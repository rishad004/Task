package database

import (
	"fmt"
	"temp/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func InitDatabse() {
	dsn := "host=localhost user=postgres password=123 dbname=week6 port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("ERROR CAN'T CONNECT TO DATABASE")
	}
	DataBase = db
	DataBase.AutoMigrate(&model.Users{}, &model.Admins{})
}
