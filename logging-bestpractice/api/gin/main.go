package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"git.rarejob.com/shintaro.ikeda/platform_logging/new_gin_logging"
	logger "git.rarejob.com/shintaro.ikeda/platform_logging/new_logger"
)

const (
	ServiceName = "Gin Sample App"
)

var (
	version  = "unknown"
	revision = "unknown"
	host, _  = os.Hostname()
	begin    = time.Now()
)

func main() {
	router := gin.New()
	{
		router.Use(
			//gin_logging.LoggerWithWriter(c.Logger.AccessLogWriter /* Exclude Path */, "/example/path"),
			new_gin_logging.LoggerWithWriter(logger.DefaultLogger() /* Exclude Path */, "/example/path"),
			Recovery(), // Recovery is a Gin middleware to handle panic error of application.
		)
	}

	v2 := router.Group("/v2")
	{
		sample := v2.Group("/sample")
		{
			sample.GET("", HandlerSampleEndpoint)
		}
	}

	logger.DefaultLogger().Info("", // No transaction ID since it's initialization
		ServiceName,
		"server is started",
		logger.F("host", host),
		// logger.F("port", c.HTTP.Port),
		logger.F("version", version),
		logger.F("revision", revision),
	)

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.DefaultLogger().Panic("", // No transaction ID since it's initialization
			ServiceName,
			"failed to listen and serve",
			logger.E(err))
	}
}
