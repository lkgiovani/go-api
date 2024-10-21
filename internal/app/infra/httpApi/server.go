package httpApi

import (
	"fmt"
	"go-api/internal/app/api/router"
	"go-api/internal/app/infra/config"
	"net/http"
)

// NewServer cria e retorna um novo mux de rotas
func NewServer() *http.ServeMux {
	mux := http.NewServeMux()

	// Inicializa as rotas
	routers := router.InitializeRoutes()
	mux.Handle("/", routers)

	return mux
}

func StartServer(mux *http.ServeMux, config *config.Config) {
	port := ":" + config.Server.Port
	fmt.Printf("Servidor iniciado na porta %s\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
