package db

import (
	"go-api/internal/app/infra/config"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"db",
	fx.Provide(NewDBConfig),
	fx.Invoke(func(d *DBConfig, config *config.Config, lc fx.Lifecycle) {
		err := d.NewDB(config)
		if err != nil {
			panic(err)
		}

		d.InitDB(lc, config)
	}),
)
