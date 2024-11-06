package db

import (
	"go.uber.org/fx"
)

var Module = fx.Module("db",
	fx.Provide(NewDBConfig), // Fornece *sql.DB diretamente
	fx.Invoke(InitDB),       // Configura o ciclo de vida do banco e inicialização
)
