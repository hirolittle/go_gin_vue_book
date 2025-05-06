package main

// 获取客户端的 IP 地址

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ip", func(c *gin.Context) {
		clientIP := c.ClientIP()
		c.JSON(http.StatusOK, gin.H{
			"message":  "Hi, this is a test",
			"clientIP": clientIP,
		})
	})

	r.Run(":8080")
}
