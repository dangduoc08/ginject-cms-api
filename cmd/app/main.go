package main

import (
	conf "github.com/dangduoc08/ginject-cms-api/internal/infrastructure/config"
	"github.com/dangduoc08/ginject-cms-api/internal/infrastructure/health"
	"github.com/dangduoc08/ginject-cms-api/internal/infrastructure/postgres"
	"github.com/dangduoc08/ginject/core"
	"github.com/dangduoc08/ginject/log"
	"github.com/dangduoc08/ginject/middlewares"
	"github.com/dangduoc08/ginject/modules/config"
	"github.com/dangduoc08/ginject/versioning"
)

func main() {
	app := core.New()

	logger := log.NewLog(&log.LogOptions{
		Level:     log.DebugLevel,
		LogFormat: log.PrettyFormat,
	})

	app.
		EnableVersioning(
			versioning.Versioning{
				Type:           versioning.HEADER,
				DefaultVersion: versioning.NEUTRAL_VERSION,
				Key:            "X-Api-Version",
			},
		).
		UseLogger(logger).
		BindGlobalMiddlewares(
			middlewares.RequestLogger{},
			middlewares.CORS{
				AllowOrigin:        conf.Env.DomainWhitelist,
				IsAllowCredentials: true,
				AllowHeaders: []string{
					"Origin",
					"X-Requested-With",
					"Content-Type",
					"Accept",
					"v",
					"Cookie",
				},
			},
		)

	app.Create(
		core.ModuleBuilder().
			Imports(
				conf.ConfModule,
				health.HealthModule,
				postgres.PostgresModule,
			).
			Build(),
	)

	configService := app.Get(config.ConfigService{}).(config.ConfigService)

	app.Logger.Fatal("AppError", "error", app.Listen(configService.Get("SERVER_PORT").(int)))
}
