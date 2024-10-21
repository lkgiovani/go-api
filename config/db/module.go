package db

import "go.uber.org/fx"

var Module = fx.Module(
	"db",
	fx.Provide(NewDB), // Fornece a instância do banco de dados
	fx.Invoke(InitDB), // Invoca a função de inicialização
)
