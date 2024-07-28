package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRouter() error {
	router := gin.Default()

	api := router.Group("/api")
	{
		//router configuration
		api.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type", "Content-Length", "X-Requested-With", "Rest-Api-Key", "Client-Id", "Client-Secret", "Client-Key"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		//version
		api.GET("/version", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{
				"version": "v1.0.0",
				"time":    time.Now().Unix(),
				"status":  "active",
			}})
		})

		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong")
		})

	}

	return router.Run(":2514")
}
