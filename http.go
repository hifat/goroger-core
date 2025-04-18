package core

import "context"

type IHttpEngine interface {
	IHttpRouter
	Listener(port string) error
}

type IHttpCtx interface {
	Context() context.Context
	Method() string
	SetHeader(key string, val string)
	Param(key string) string
	BodyParser(v any) error
	SendString(message string)
	JSON(status int, v any)
	Abort(status int, message string)
	AbortWithJSON(status int, v any)
	Set(key string, value any)
	Get(key string) any
	MustGet(key string) (value any, exists bool)
	QueryParser(v any) error
	Next()
}

type IHttpRouter interface {
	Use(handlers ...func(IHttpCtx)) IHttpEngine
	Group(prefix string, handlers ...func(IHttpCtx)) IHttpRouter
	Get(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Post(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Put(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Patch(path string, handlers ...func(IHttpCtx)) IHttpRouter
	Delete(path string, handlers ...func(IHttpCtx)) IHttpRouter
}
