package framework

import (
	"github.com/gofiber/fiber/v2"
	core "github.com/hifat/goroger-core"
)

type httpEngine struct {
	engine *fiber.App
}

func NewFiberEngineCtx(engine *fiber.App) core.IHttpEngine {
	return &httpEngine{engine}
}

func (e *httpEngine) Use(handlers ...func(core.IHttpCtx)) core.IHttpEngine {
	fiberHandlers := make([]interface{}, len(handlers))
	for i, handler := range handlers {
		fiberHandlers[i] = func(ctx *fiber.Ctx) error {
			handler(NewFiberCtx(ctx))
			return nil
		}
	}

	e.engine.Use(fiberHandlers...)

	return e
}

func (e *httpEngine) Get(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		fiberHandlers[i] = func(ctx *fiber.Ctx) error {
			handler(NewFiberCtx(ctx))
			return nil
		}
	}

	e.engine.Get(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Post(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		fiberHandlers[i] = func(ctx *fiber.Ctx) error {
			handler(NewFiberCtx(ctx))
			return nil
		}
	}

	e.engine.Post(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Put(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		fiberHandlers[i] = func(ctx *fiber.Ctx) error {
			handler(NewFiberCtx(ctx))
			return nil
		}
	}

	e.engine.Put(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Patch(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		fiberHandlers[i] = func(ctx *fiber.Ctx) error {
			handler(NewFiberCtx(ctx))
			return nil
		}
	}

	e.engine.Patch(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Delete(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		fiberHandlers[i] = func(ctx *fiber.Ctx) error {
			handler(NewFiberCtx(ctx))
			return nil
		}
	}

	e.engine.Delete(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Listener(port string) error {
	return e.engine.Listen(port)
}
