package server

import (
	"api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {

	// sample handler
	sampleHandler := &handler.SampleHandler{}
	r.GET("/sample", authMiddleware(), sampleHandler.Index)

	r.GET("/health-check", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "running"})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// authorization
		isAuthorized := true
		if !isAuthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized."})
			c.Abort()
		}
	}
}
