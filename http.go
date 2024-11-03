package core

type IHttpEngine interface {
	IHttpRouter
	Listener(port string) error
}

type IHttpCtx interface {
	Param(key string) string
	ShouldBind(v any) error
	SendString(message string)
	JSON(status int, v any) error
	Abort(status int, message string)
	Set(key string, value any)
	Get(key string) any
	MustGet(key string) (value any, exists bool)
	ShouldBindQuery(v any) error
}

type IHttpRouter interface {
	Use(handlers ...func(IHttpCtx)) IHttpEngine
	Get(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Post(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Put(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Patch(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Delete(path string, handlers ...func(IHttpCtx)) IHttpRouter
}
