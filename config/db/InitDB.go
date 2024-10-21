package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-api/internal/app/infra/config"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
)

var DB *sql.DB

func NewDB(config *config.Config) (*sql.DB, error) {

	dns := config.Mysql.Url
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error checking the connection to the database: %v", err)
	}

	fmt.Println("passo aqui ")
	return db, nil
}

func InitDB(lc fx.Lifecycle, config *config.Config, db *sql.DB) {
	databaseName := "BDTest"

	fmt.Println("passo quie 1")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("OnStart do InitDB foi chamado")

			err := createDatabase(databaseName)
			if err != nil {
				return errors.New("Error creating the database: " + err.Error())
			}

			dbNameDSN := fmt.Sprintf(config.Mysql.Url + databaseName)
			DB, err = sql.Open("mysql", dbNameDSN)
			if err != nil {
				return err
			}

			err = DB.Ping()
			if err != nil {
				return err
			}
			fmt.Println("passo aqui ")
			sqlFilePath := "config/db/migrations/migrationInit.sql"
			return executeSQLFile(sqlFilePath)

		},
		OnStop: func(ctx context.Context) error {
			return DB.Close()
		},
	})

	fmt.Println("passo aqui 2")

}

// Function to create the database if it doesn't exist
func createDatabase(name string) error {
	fmt.Println("Tentando criar o banco de dados:", name)
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	_, err := DB.Exec(query)
	if err != nil {
		return errors.New("Error creating the database: " + err.Error())
	}

	return nil
}

// Function to read and execute the SQL file
func executeSQLFile(filePath string) error {
	// Read the SQL file
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New("Error reading the SQL file: " + err.Error())
	}

	// Convert the file content to string
	sqlString := string(sqlBytes)

	// Execute the SQL in the database
	_, err = DB.Exec(sqlString)
	if err != nil {
		return errors.New("Error executing the SQL file: " + err.Error())
	}

	return nil
}
