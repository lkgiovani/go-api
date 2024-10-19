package httpApi

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"httpapi",
	fx.Provide(NewServer),
	fx.Invoke(StartServer),
)
