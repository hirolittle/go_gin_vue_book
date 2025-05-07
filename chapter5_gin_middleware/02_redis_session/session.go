package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func authMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("user_id", userId)
	c.Next()
}

func main() {

	router := gin.Default()
	store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", "", []byte("secret"))
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return
	}
	fmt.Printf("store: %v\n", store)

	router.Use(sessions.Sessions("mySession02", store))

	router.GET("/login", func(c *gin.Context) {
		// 检查用户名和密码
		validUser := true
		if validUser {
			session := sessions.Default(c)
			session.Set("user_id", "88")
			session.Save()
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
		}
	})

	authorized := router.Group("/auth")
	authorized.Use(authMiddleware)
	{
		authorized.GET("/protected", func(c *gin.Context) {
			userId, exists := c.Get("user_id")
			if exists {
				c.JSON(http.StatusOK, gin.H{
					"message": "Hello, " + userId.(string),
				})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			}
		})
	}

	router.Run(":8080")

}
