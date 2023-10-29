package gin

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinContext struct {
	*gin.Context
}

func NewGinContext(c *gin.Context) *GinContext {
	return &GinContext{Context: c}
}

// implements restapi.Context
func (c *GinContext) JSON(code int, obj interface{}) {
	c.Context.JSON(code, obj)
}

func (c *GinContext) Bind(obj interface{}) error {
	return c.Context.ShouldBindJSON(obj)
}

func (c *GinContext) GetQuery(key string) string {
	return c.Context.Query(key)
}

func (c *GinContext) Set(key string, value any) {
	c.Context.Set(key, value)
}

func (c *GinContext) Get(key string) (value any, exists bool) {
	return c.Context.Get(key)
}

func (c *GinContext) GetParam(key string) (int, error) {
	param := c.Context.Param(key)

	v, err := strconv.Atoi(param)
	if err != nil {
		return 0, err

	}

	return v, nil
}

func (c *GinContext) AttachError(err error) *gin.Error {
	return c.Context.Error(err)
}
