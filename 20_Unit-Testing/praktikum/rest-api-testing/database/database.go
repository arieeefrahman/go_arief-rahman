package database

import (
	"fmt"
	"log"
	"prak-testing/config"
	"prak-testing/models"
	"prak-testing/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	DB, err = gorm.Open(mysql.Open(connectionString))

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	DB.AutoMigrate(&models.User{}, &models.Book{})
}

func InitTestDB() {
	c := config.InitConfig()
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DB_Username,
		c.DB_Password,
		c.DB_Host,
		c.DB_Port,
		util.GetConfig("DB_TEST_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString))
	
	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	DB.AutoMigrate(&models.User{}, &models.Book{})
}

func SeedUser() models.User {
	var user models.User = models.User{
		Name: "test",
		Email: "test@gmail.com",
		Password: "test123",
	}

	if err := DB.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser models.User

	DB.Last(&createdUser)

	return createdUser
}

func SeedBook() models.Book {
	var book models.Book = models.Book{
		Title: "test title",
		Isbn: "123456",
		Description: "test title description",
	}

	if err := DB.Create(&book).Error; err != nil {
		panic(err)
	}

	var createdBook models.Book

	DB.Last(&createdBook)

	return createdBook
}