// Package db provides a database configuration and initialization functionality.
package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-api/internal/app/infra/config/configEnv"
	"go-api/pkg/projectError"
	"go.uber.org/fx"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NewDBConfig establishes a new database connection based on the provided configuration.
func NewDBConfig(config *configEnv.Config) (*sql.DB, error) {
	// Create a DSN (Data Source Name) from the configuration.
	dns := config.Mysql.Url
	// Open a database connection using the DSN.
	db, err := sql.Open("mysql", dns)
	if err != nil {
		// Return an error if the connection fails.
		return nil, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to connect to database",
			Path:      "config/db/InitDB.go",
			PrevError: err,
		}
	}

	// Ping the database to verify the connection.
	err = db.Ping()
	if err != nil {
		// Return an error if the ping fails.
		return nil, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to ping database",
			Path:      "config/db/InitDB.go",
			PrevError: err,
		}
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso")
	return db, nil
}

// InitDB initializes the database using the provided lifecycle and configuration.
func InitDB(lc fx.Lifecycle, db *sql.DB, config *configEnv.Config) error {
	// Define o ciclo de vida do banco de dados no Fx
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { // Adiciona o contexto como parâmetro

			fmt.Println("Iniciando banco de dados...")

			// Define o nome do banco de dados
			databaseName := "BDTest"

			// Cria o banco de dados, se não existir
			err := createDatabase(db, databaseName)
			if err != nil {
				return &projectError.Error{
					Code:      projectError.EINTERNAL,
					Message:   "failed to create database",
					Path:      "config/db/InitDB.go",
					PrevError: err,
				}
			}

			// Cria o DSN para o banco de dados
			dbNameDSN := fmt.Sprintf(config.Mysql.Url + databaseName)
			// Abre a conexão com o banco de dados usando o DSN
			db, err = sql.Open("mysql", dbNameDSN)
			if err != nil {
				return &projectError.Error{
					Code:      projectError.EINTERNAL,
					Message:   "failed to connect to database",
					Path:      "config/db/InitDB.go",
					PrevError: err,
				}
			}

			// Verifica a conexão com o banco de dados
			err = db.Ping()
			if err != nil {
				return &projectError.Error{
					Code:      projectError.EINTERNAL,
					Message:   "failed to ping database",
					Path:      "config/db/InitDB.go",
					PrevError: err,
				}
			}

			// Executa o arquivo SQL para inicializar o banco de dados
			sqlFilePath := "internal/app/infra/config/db/migrations/migrationInit.sql"
			err = executeSQLFile(db, sqlFilePath)
			if err != nil {
				return &projectError.Error{
					Code:      projectError.EINTERNAL,
					Message:   "failed to execute SQL file",
					Path:      "config/db/InitDB.go",
					PrevError: err,
				}
			}
			return nil
		},
		OnStop: func(ctx context.Context) error { // Adiciona o contexto como parâmetro
			fmt.Println("Fechando conexão com o banco de dados...")
			// Fecha a conexão com o banco de dados
			return db.Close()
		},
	})

	return nil
}

// createDatabase creates the database if it doesn't exist.
func createDatabase(d *sql.DB, name string) error {
	// Create a query to create the database.
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	// Execute the query using the database connection.
	_, err := d.Exec(query) // Use the db passed as a parameter
	if err != nil {
		// Return an error if the query execution fails.
		return &projectError.Error{
			Code:      projectError.ENOTFOUND,
			Message:   "Database not found. " + err.Error(),
			Path:      "config/db/InitDB.go",
			PrevError: err,
		}
	}

	return nil
}

// executeSQLFile reads and executes the SQL file at the given file path.
// It returns an error if the file cannot be read or the SQL execution fails.
func executeSQLFile(db *sql.DB, filePath string) error {

	// Obtenha o diretório de trabalho atual
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Construa o caminho completo
	fullPath := filepath.Join(workingDir, filePath)

	sqlBytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.ENOTFOUND,
			Message:   "failed to read file: " + err.Error(),
			Path:      "config/db/InitDB.go",
			PrevError: err,
		}
	}

	sqlString := string(sqlBytes)
	_, err = db.Exec(sqlString)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.ENOTFOUND,
			Message:   "failed to execute SQL: " + err.Error(),
			Path:      "config/db/InitDB.go",
			PrevError: err,
		}
	}

	return nil
}
