package restapi

type defaultRouter interface {
	GET(relativePath string, handlers ...interface{})
	POST(relativePath string, handlers ...interface{})
	PUT(relativePath string, handlers ...interface{})
	DELETE(relativePath string, handlers ...interface{})
}

type Router interface {
	defaultRouter
	Group(relativePath string, handlers ...interface{}) RouterGroup
}

type RouterGroup interface {
	defaultRouter
}
