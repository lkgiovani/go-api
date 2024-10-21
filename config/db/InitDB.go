// Package db provides a database configuration and initialization functionality.
package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-api/internal/app/infra/config"
	"go-api/pkg/projectError"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
)

// DBConfig represents a database configuration.
type DBConfig struct {
	// DB is the database connection.
	DB *sql.DB
}

// NewDBConfig returns a new instance of DBConfig.
func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

// NewDB establishes a new database connection based on the provided configuration.
func (d *DBConfig) NewDB(config *config.Config) error {
	// Create a DSN (Data Source Name) from the configuration.
	dns := config.Mysql.Url
	// Open a database connection using the DSN.
	db, err := sql.Open("mysql", dns)
	if err != nil {
		// Return an error if the connection fails.
		return fmt.Errorf("Error connecting to database: %v", err)
	}

	// Ping the database to verify the connection.
	err = db.Ping()
	if err != nil {
		// Return an error if the ping fails.
		return fmt.Errorf("Error checking the connection to the database: %v", err)
	}

	// Print a message to indicate the connection was successful.
	fmt.Println("passo aqui ")
	// Store the database connection in the DBConfig instance.
	d.DB = db

	return nil
}

// InitDB initializes the database using the provided lifecycle and configuration.
func (d *DBConfig) InitDB(lc fx.Lifecycle, config *config.Config) {
	// Define the database name.
	databaseName := "BDTest"

	// Print a message to indicate the initialization process has started.
	fmt.Println("passo quie 1")

	// Append a hook to the lifecycle to execute on start and stop.
	lc.Append(fx.Hook{
		// OnStart is called when the application starts.
		OnStart: func(ctx context.Context) error {
			// Print a message to indicate the OnStart hook has been called.
			log.Println("OnStart do InitDB foi chamado")

			// Create the database if it doesn't exist.
			err := d.createDatabase(databaseName)
			if err != nil {
				// Return an error if the database creation fails.
				return &projectError.Error{Code: projectError.ENOTFOUND, Message: "ProviderAuth not found."}
			}

			// Create a DSN for the database.
			dbNameDSN := fmt.Sprintf(config.Mysql.Url + databaseName)
			// Open a database connection using the DSN.
			d.DB, err = sql.Open("mysql", dbNameDSN)
			if err != nil {
				// Return an error if the connection fails.
				return err
			}

			// Ping the database to verify the connection.
			err = d.DB.Ping()
			if err != nil {
				// Return an error if the ping fails.
				return err
			}

			// Print a message to indicate the connection was successful.
			fmt.Println("passo aqui ")
			// Execute the SQL file to initialize the database.
			sqlFilePath := "config/db/migrations/migrationInit.sql"
			return d.executeSQLFile(sqlFilePath)

		},
		// OnStop is called when the application stops.
		OnStop: func(ctx context.Context) error {
			// Close the database connection.
			return d.DB.Close()
		},
	})

	// Print a message to indicate the initialization process has completed.
	fmt.Println("passo aqui 2")

}

// createDatabase creates the database if it doesn't exist.
func (d *DBConfig) createDatabase(name string) error {
	// Print a message to indicate the database creation process has started.
	fmt.Println("Tentando criar o banco de dados:", name)
	// Create a query to create the database.
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	// Execute the query using the database connection.
	_, err := d.DB.Exec(query) // Use the db passed as a parameter
	if err != nil {
		// Return an error if the query execution fails.
		return &projectError.Error{
			Code:    projectError.ENOTFOUND,
			Message: "Database not found. " + err.Error(),
		}
	}

	return nil
}

// executeSQLFile reads and executes the SQL file at the given file path.
// It returns an error if the file cannot be read or the SQL execution fails.
func (d *DBConfig) executeSQLFile(filePath string) error {
	// Read the SQL file
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &projectError.Error{
			Code:    projectError.ENOTFOUND,
			Message: "failed to read file: " + err.Error(),
		}
	}

	// Convert the file content to string
	sqlString := string(sqlBytes)

	// Execute the SQL in the database
	_, err = d.DB.Exec(sqlString)
	if err != nil {
		return &projectError.Error{
			Code:    projectError.ENOTFOUND,
			Message: "failed to execute SQL: " + err.Error(),
		}
	}

	return nil
}
