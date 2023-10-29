package restapi

import "github.com/gin-gonic/gin"

type Context interface {
	Bind(interface{}) error
	JSON(code int, obj interface{})
	GetQuery(string) string
	GetParam(string) (int, error)
	AttachError(err error) *gin.Error
	Set(key string, value any)
	Get(key string) (value any, exists bool)
}
