package cfg

import (
	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	// App flags
	&cli.StringFlag{
		Name:        "app-env",
		Destination: &App.Environment,
		EnvVar:      "APP_ENV",
		Value:       Development,
	},
	&cli.StringFlag{
		Name:        "app-log-level",
		Destination: &App.LogLevel,
		EnvVar:      "APP_LOG_LEVEL",
		Value:       "debug",
	},

	// Free Wheather API flags
	&cli.StringFlag{
		Name:        "freeweather-api-key",
		Destination: &FreeWeather.ApiKey,
		EnvVar:      "FREEWEATHER_API_KEY",
	},

	// OpenTelemetry flags
	&cli.StringFlag{
		Name:        "otl-endpoint",
		Destination: &Otl.OTELPEndpoint,
		EnvVar:      "OTEL_ENDPOINT",
		Value:       "localhost:4317",
	},
	&cli.StringFlag{
		Name:        "otl-service-name",
		Destination: &Otl.ServiceName,
		EnvVar:      "OTEL_SERVICE_NAME",
		Value:       "api2",
	},
}
