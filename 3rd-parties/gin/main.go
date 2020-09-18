package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type MyStruct struct {
	Key string `json:"key" binding:"required"`
}

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var h MyStruct

		if err := c.ShouldBindBodyWith(&h, binding.JSON); err != nil {
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("%+v", err),
			})
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(MyMiddleware())
	r.POST("/post", func(c *gin.Context) {
		var h MyStruct

		// [My Note]
		// このDecodeをする時点で 2重でデコードすると EOFが発生。
		// この事象はginに依るところではない。
		// c.BindJSON(&h) この中でjson.NewDecoderをしているのでEOFのエラーが発生していた。
		// c.ShouldBindBodyWith を利用すれば c.Request.Body が消化されずに済む。

		decoder := json.NewDecoder(c.Request.Body)
		err := decoder.Decode(&h)
		if err != nil {
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("%+v", err),
			})
			return
		}

		//if err := c.BindJSON(&h); err != nil {
		//	fmt.Println(err)
		//	c.JSON(400, gin.H{
		//		"message": fmt.Sprintf("err is %+v\n", err),
		//	})
		//	return
		//}

		//if err := c.ShouldBindBodyWith(&h, binding.JSON); err != nil {
		//	fmt.Println(err)
		//	c.JSON(400, gin.H{
		//		"message": fmt.Sprintf("err is %+v\n", err),
		//	})
		//	return
		//}

		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
