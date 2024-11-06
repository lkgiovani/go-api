package httpApi

import (
	"context"
	"database/sql"
	"fmt"
	"go-api/internal/app/api/router"
	"go-api/internal/app/infra/config/configEnv"
	"go.uber.org/fx"
	"net/http"
)

// NewServer cria e retorna um novo mux de rotas, injetando a dependência *sql.DB
//
// Essa função é injetada no StartServer e é usada para criar o mux de rotas
func NewServer(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Inicializa as rotas usando o banco de dados
	routers := router.InitializeRoutes(db)
	mux.Handle("/", routers)

	return mux
}

// StartServer inicializa o servidor HTTP
//
// Essa função é usada pelo Fx para inicializar o servidor HTTP
// e garantir que o servidor seja iniciado e parado corretamente
func StartServer(lc fx.Lifecycle, mux *http.ServeMux, config *configEnv.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			port := ":" + config.Server.Port
			fmt.Printf("Servidor iniciado na porta %s\n", port)
			go func() {
				if err := http.ListenAndServe(port, mux); err != nil {
					fmt.Println("Erro ao iniciar o servidor:", err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Parando o servidor HTTP...")
			// Adicione lógica de parada aqui, se necessário
			return nil
		},
	})
}
