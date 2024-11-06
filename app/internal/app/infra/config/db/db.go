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
	dns := config.Mysql.Url + "bdtest"
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

func InitDB(lc fx.Lifecycle, config *configEnv.Config) (*sql.DB, error) { // Passo 1: Conexão inicial ao MySQL sem especificar o banco// Passo 1: Conexão inicial ao MySQL sem especificar o banco
	db, err := sql.Open("mysql", config.Mysql.Url)
	if err != nil {
		return nil, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to connect to MySQL server",
			Path:      "config/db/InitDB.go",
			PrevError: err,
		}
	}
	defer db.Close() // Fecha esta conexão inicial ao final

	// Passo 2: Criação do banco de dados, se necessário
	databaseName := "bdtest"
	err = createDatabase(db, databaseName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Banco de dados", databaseName, "criado com sucesso ou já existia.")

	// Passo 3: Conectar-se com o banco de dados `bdtest`
	db, err = sql.Open("mysql", config.Mysql.Url+"bdtest")
	if err != nil {
		return nil, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to connect to database",
			Path:      "config/db/InitDB.go",
			PrevError: err,
		}
	}

	err = db.Ping()
	if err != nil {
		return nil, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to ping database",
			Path:      "config/db/InitDB.go",
			PrevError: err,
		}
	}

	// Passo 4: Executar o arquivo SQL de migração
	sqlFilePath := "internal/app/infra/config/db/migrations/migrationInit.sql"
	err = executeSQLFile(db, sqlFilePath)
	if err != nil {
		return nil, err
	}
	fmt.Println("Arquivo de migração SQL executado com sucesso.")

	// Configuração do ciclo de vida com Fx
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			fmt.Println("Fechando conexão com o banco de dados...")
			return db.Close()
		},
	})

	fmt.Println("Conexão com o banco de dados configurada e pronta.")
	return db, nil
}

// createDatabase creates the database if it doesn't exist.
func createDatabase(d *sql.DB, name string) error {
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	_, err := d.Exec(query)
	if err != nil {
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
func executeSQLFile(db *sql.DB, filePath string) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

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
