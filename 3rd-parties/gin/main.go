package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		b := c.Request.Body
		fmt.Println("in middleware", b)

		c.Next()

		// after request
	}
}

type Header struct {
	Value string `json:"value" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Use(MyMiddleware())
	r.POST("/post", func(c *gin.Context) {

		b := c.Request.Body
		fmt.Println("in post", b)
		defer b.Close()

		var h Header
		// h.Value = "value"

		if err := c.BindJSON(&h); err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("err is %+v\n", err),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
