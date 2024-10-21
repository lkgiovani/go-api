package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-api/config/db"
	"go-api/internal/app/infra/config"
	"go-api/internal/app/infra/httpApi"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error em het .env")
	}

	app := fx.New(
		fx.WithLogger(
			func() fxevent.Logger {
				return &fxevent.ConsoleLogger{W: log.Writer()}
			},
		),
		config.Module,
		db.Module,      // O módulo de banco de dados será iniciado primeiro
		httpApi.Module, // O módulo HTTP será iniciado depois
	)

	app.Run()
	// Inicializa o banco de dados

	fmt.Println("Aplicação encerrada.")

}
