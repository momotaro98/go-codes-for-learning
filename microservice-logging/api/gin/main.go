package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	logger "github.com/momotaro98/go-codes-for-learning/microservice-logging/logger"
)

const (
	ServiceName = "Gin Sample App"
)

var (
	version  = "unknown"
	revision = "unknown"
	host, _  = os.Hostname()
)

func main() {
	router := gin.New()
	{
		router.Use(
			LoggerWithWriter(logger.DefaultLogger() /* Exclude Path */, "/example/path"),
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
