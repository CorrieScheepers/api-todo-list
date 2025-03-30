package database

import (
	"fmt"
	"log"
	"os"

	"api-todo-list/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectToMySqlDb() (*gorm.DB, error) {
	err := godotenv.Load(".config/.config")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDb := os.Getenv("MYSQL_DB")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDb)

	database, err := gorm.Open("mysql", dataSource)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check if the connection is successful
	if err := database.DB().Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to MySQL database successfully")

	database.AutoMigrate(&models.TaskModel{})
	return database, nil
}
