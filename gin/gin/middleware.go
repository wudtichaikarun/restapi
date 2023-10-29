package gin

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wudtichaikarun/restapi/response/response"
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

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, &response.FailResponse{
					Code:    http.StatusInternalServerError,
					Status:  response.FailStatus,
					Message: "Internal Server Error",
				})
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if responseFailed, ok := err.Err.(*response.FailResponse); ok {
					c.JSON(responseFailed.Code, &response.FailResponse{
						Code:    responseFailed.Code,
						Status:  responseFailed.Status,
						Message: responseFailed.Error(),
					})
				} else {
					c.JSON(http.StatusInternalServerError, &response.FailResponse{
						Code:    http.StatusInternalServerError,
						Status:  response.FailStatus,
						Message: err.Err.Error(),
					})
				}
			}

			c.Abort()
		}
	}
}
