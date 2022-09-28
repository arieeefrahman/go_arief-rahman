package database

import (
	"fmt"
	"prak-18-prob-2/config"
	"prak-18-prob-2/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	config := config.InitConfig()
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	// DB.AutoMigrate(&User{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}