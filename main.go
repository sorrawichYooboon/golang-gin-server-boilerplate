package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("ping", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Pong!",
		// })
		c.String(200, "Pong!")
	})

	router.Run(":8080")
}
