package db

import (
	"fmt"
	"log"
	c "weekly-task-2/config"
	"weekly-task-2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func InitDB() {
	cfg := c.InitConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", 
		cfg.DB_Username, 
		cfg.DB_Password,
		cfg.DB_Host,
		cfg.DB_Port,
		cfg.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	DB.AutoMigrate(&models.Category{}, &models.Product{})
}