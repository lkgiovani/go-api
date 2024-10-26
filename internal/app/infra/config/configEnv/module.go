package configEnv

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewConfig,
)
