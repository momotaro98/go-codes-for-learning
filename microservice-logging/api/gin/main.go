package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/momotaro98/go-codes-for-learning/microservice-logging/logger"
)

func main() {
	router := gin.New()
	{
		router.Use(
			LoggerWithWriter(logger.Log /* Exclude Path */, "/example/path"),
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

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Log.Error("", // No transaction ID since it's initialization
			"failed to listen and serve",
			logger.E(err))
	}
}
