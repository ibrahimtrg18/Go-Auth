package config

import (
	"fmt"

	model "github.com/ibrahimtrg18/jemari/cmd/jemari/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDatabase() {
	environment := GetEnvironment()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", environment.User, environment.Password, environment.Host, environment.Port, environment.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.PartnerUser{})
}

func InjectDatabase() *gorm.DB {
	return db
}
