package gin

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wudtichaikarun/restapi"
)

type GinRouter struct {
	engine *gin.Engine
}

type GinRouterGroup struct {
	engine *gin.RouterGroup
}

// implements restapi.Router.
func (r *GinRouter) GET(relativePath string, handlers ...interface{}) {
	r.engine.GET(relativePath, convertHandlersToGin(handlers)...)
}

// implements restapi.Router.
func (r *GinRouter) POST(relativePath string, handlers ...interface{}) {
	r.engine.POST(relativePath, convertHandlersToGin(handlers)...)
}

// implements restapi.Router.
func (r *GinRouter) PUT(relativePath string, handlers ...interface{}) {
	r.engine.PUT(relativePath, convertHandlersToGin(handlers)...)
}

// implements restapi.Router.
func (r *GinRouter) DELETE(relativePath string, handlers ...interface{}) {
	r.engine.DELETE(relativePath, convertHandlersToGin(handlers)...)
}

// implements restapi.Router.
func (r *GinRouter) Group(relativePath string, handlers ...interface{}) restapi.RouterGroup {
	group := r.engine.Group(relativePath, convertHandlersToGin(handlers)...)
	return &GinRouterGroup{engine: group}
}

func NewGinRouter() *GinRouter {
	r := gin.Default()

	r.Use(CORS())
	r.Use(ErrorHandler())

	// initial swagger doc
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &GinRouter{engine: r}
}

func convertHandlersToGin(handlers []interface{}) []gin.HandlerFunc {
	var ginHandlers []gin.HandlerFunc

	for _, h := range handlers {

		if handlerFunc, ok := h.(func(c *gin.Context)); ok {
			ginHandlers = append(ginHandlers, gin.HandlerFunc(handlerFunc))
		} else if ginHandler, ok := h.(gin.HandlerFunc); ok {
			ginHandlers = append(ginHandlers, ginHandler)
		} else if handlerFunc, ok := h.(func(c restapi.Context)); ok {
			ginHandlers = append(ginHandlers, convertToGinHandler(handlerFunc))
		} else {
			panic("unimplemented")
		}
	}

	return ginHandlers
}

func convertToGinHandler(handler func(restapi.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&GinContext{Context: c})
	}
}

// GET implements restapi.RouterGroup.
func (rg *GinRouterGroup) GET(relativePath string, handlers ...interface{}) {
	rg.engine.GET(relativePath, convertHandlersToGin(handlers)...)
}

// DELETE implements restapi.RouterGroup.
func (rg *GinRouterGroup) DELETE(relativePath string, handlers ...interface{}) {
	rg.engine.DELETE(relativePath, convertHandlersToGin(handlers)...)
}

// POST implements restapi.RouterGroup.
func (rg *GinRouterGroup) POST(relativePath string, handlers ...interface{}) {
	rg.engine.POST(relativePath, convertHandlersToGin(handlers)...)
}

// PUT implements restapi.RouterGroup.
func (rg *GinRouterGroup) PUT(relativePath string, handlers ...interface{}) {
	rg.engine.PUT(relativePath, convertHandlersToGin(handlers)...)
}
