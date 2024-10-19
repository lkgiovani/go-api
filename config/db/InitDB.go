package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(databaseName string) {
	var err error

	dsn := "root:root@tcp(127.0.0.1:3306)/"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Check the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error checking the connection to the database: %v", err)
	}
	fmt.Println("Connection to the database established successfully!")

	createDatabase(databaseName)

	// Now connect to the newly created database
	DB, err = sql.Open("mysql", fmt.Sprintf("root:root@tcp(localhost:3306)/%s", databaseName))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Check the connection to the specific database
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error checking the connection to the database: %v", err)
	}
	fmt.Printf("Connection to the database '%s' established successfully!\n", databaseName)

	sqlFilePath := "config/db/migrations/migrationInit.sql"
	executeSQLFile(sqlFilePath)
}

// Function to create the database if it doesn't exist
func createDatabase(name string) {
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating the database '%s': %v", name, err)
	}
	fmt.Printf("Database '%s' created or already exists.\n", name)
}

// Function to read and execute the SQL file
func executeSQLFile(filePath string) {
	// Read the SQL file
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading the SQL file: %v", err)
	}

	// Convert the file content to string
	sqlString := string(sqlBytes)

	// Execute the SQL in the database
	_, err = DB.Exec(sqlString)
	if err != nil {
		log.Fatalf("Error executing the SQL: %v", err)
	}

	fmt.Println("SQL file executed successfully!")
}
