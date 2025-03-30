package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func connectToMySqlDb() (*sql.DB, error) {
	dataSource := "user:password@tcp(localhost:3316)/todo_list_db"

	database, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check if the connection is successful
	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to MySQL database successfully")
	return database, nil
}
