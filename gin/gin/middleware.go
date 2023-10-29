package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()

	// You can customize the CORS configuration here
	// For example, to allow only specific origins:
	// config.AllowOrigins = []string{"http://example.com"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	return cors.New(config)
}
