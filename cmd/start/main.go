package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-api/config/db"
	"go-api/internal/app/infra/httpApi"
	"go.uber.org/fx"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error em het .env")
	}

	databaseName := "BDTest"
	db.InitDB(databaseName)

	app := fx.New(
		httpApi.Module,
	)

	app.Run()
	// Inicializa o banco de dados

}
