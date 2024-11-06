package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-api/internal/app/infra/config/configEnv"
	"go-api/internal/app/infra/config/db"
	"go-api/internal/app/infra/httpApi"
	"go.uber.org/fx"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error em het .env")
	}

	app := fx.New(
		configEnv.Module, // Configuração do ambiente
		db.Module,        // Módulo de banco de dados, que inclui o ciclo de vida com InitDB
		httpApi.Module,   // Módulo HTTP, que depende do banco de dados
	)

	app.Run()
	// Inicializa o banco de dados

	fmt.Println("Aplicação encerrada.")

}
