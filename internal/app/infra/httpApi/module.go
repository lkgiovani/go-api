package httpApi

import (
	"go.uber.org/fx"
)

// Módulo do HTTP API
var Module = fx.Module(
	"httpapi",
	fx.Provide(NewServer),  // Fornece o *http.ServeMux
	fx.Invoke(StartServer), // Invoca o StartServer (que usa o *http.ServeMux e configEnv.Config)
)
