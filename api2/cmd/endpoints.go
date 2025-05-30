package cmd

import (
	"fmt"

	_ "github.com/HenriqueOtsuka/cep-weather/api2/docs"
	"github.com/HenriqueOtsuka/cep-weather/api2/internal/cfg"
	"github.com/HenriqueOtsuka/cep-weather/api2/internal/controllers/app"
	"github.com/HenriqueOtsuka/cep-weather/api2/internal/controllers/cep"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// @title CEP API
// @version 1
// @description This is the CEP API
// @contact.name   Hydde
// @contact.url    https://www.github.com/HenriqueOtsuka
// @tag.name App
func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(otelgin.Middleware(cfg.Otl.ServiceName))
	middleware := NewMiddleware()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middleware.Errors)
	router.Use(middleware.Options)
	router.NoRoute(middleware.NotFound)
	router.NoMethod(middleware.MethodNotAllowed)

	router.GET("/appstatus", app.HandleGetAppStatus)
	router.GET("/cep/:cep", cep.HandleGetCepTemperature)

	return router
}
