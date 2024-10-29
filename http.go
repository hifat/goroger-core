package core

type IHttpEngine interface {
	Get(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Listener(port string) error
}

type IHttpCtx interface {
	Get(key string) string
}

type IHttpRouter interface{}
