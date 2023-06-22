package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = false
	config.AllowOrigins = []string{os.Getenv("ALLOW_ORIGIN")}
	return cors.New(config)
}
