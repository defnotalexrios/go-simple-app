package main

import (
	"{{.Module}}/app/service/api/v1"
	"{{.Module}}/extensions/xfx"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			zap.NewExample,
			xfx.AsRoute(v1.NewEchoHandler),
			xfx.AsRoute(v1.NewHelloHandler),
			fx.Annotate(
				xfx.NewServeMux,
				xfx.RouterAnnotation(),
			)),
		fx.Invoke(xfx.NewHTTPServer),
	).Run()
}
