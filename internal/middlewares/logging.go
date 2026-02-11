package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Antes de la petición (Start time)
		t := time.Now()

		// 2. Procesar la petición
		c.Next()

		// 3. Después de la petición (End time)
		latency := time.Since(t)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		log.Printf("[API-LOG] %s | %d | %v | %s", method, status, latency, path)
	}
}
