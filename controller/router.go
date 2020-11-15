package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/niwek/kwl-forwardauth/config"
)

// Router ...
var Router router

type router struct{}

// Setup the Gin Router
func (r router) Setup() *gin.Engine {
	router := gin.Default()

	corsMiddleware := cors.New(cors.Config{
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowMethods:     []string{"GET", "OPTIONS", "POST", "PUT", "DELETE", "PATCH"},
		MaxAge:           12 * time.Hour,
		AllowOrigins:     strings.Split(config.Env.AllowedOrigins, ","),
		AllowCredentials: true,
	})

	router.Use(corsMiddleware)

	health := router.Group("/health")
	{
		health.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	v1 := router.Group("/api/v1")

	user := v1.Group("user")
	{
		user.POST("", CreateUser)
		user.GET("/:id", GetUserByID)
	}

	return router
}
