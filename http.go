package core

type IHttpEngine interface {
	Use(handlers ...func(IHttpCtx)) IHttpEngine
	Get(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Post(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Put(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Patch(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Delete(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Listener(port string) error
}

type IHttpCtx interface {
	Get(key string) string
	ShouldBind(v interface{}) error
}

type IHttpRouter interface{}
