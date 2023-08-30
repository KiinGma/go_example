package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ip", func(c *gin.Context) {
		ip := c.ClientIP()
		c.String(200, "IP Address: %s", ip)
	})

	r.Run(":8080")
}
