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

func toFiberHandler(handlers []func(core.IHttpCtx)) []fiber.Handler {
	fiberHandlers := make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		fiberHandlers[i] = func(ctx *fiber.Ctx) error {
			handler(NewFiberCtx(ctx))
			return nil
		}
	}

	return fiberHandlers
}

func (e *httpEngine) Use(handlers ...func(core.IHttpCtx)) core.IHttpEngine {
	fiberHandlers := toFiberHandler(handlers)
	e.engine.Use(fiberHandlers)

	return e
}

func (e *httpEngine) Group(prefix string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)

	return &fiberGroup{
		engine: e.engine.Group(prefix, fiberHandlers...),
	}
}

func (e *httpEngine) Get(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)
	e.engine.Get(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Post(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)
	e.engine.Post(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Put(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)
	e.engine.Put(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Patch(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)
	e.engine.Patch(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Delete(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)
	e.engine.Delete(path, fiberHandlers...)

	return e
}

func (e *httpEngine) Listener(port string) error {
	return e.engine.Listen(port)
}

/* -------------------------------------------------------------------------- */
/*                                 Fiber Group                                */
/* -------------------------------------------------------------------------- */

type fiberGroup struct {
	engine fiber.Router
}

func (e *fiberGroup) Group(prefix string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)

	return &fiberGroup{
		engine: e.engine.Group(prefix, fiberHandlers...),
	}
}

func (e *fiberGroup) Use(handlers ...func(core.IHttpCtx)) core.IHttpEngine {
	fiberHandlers := toFiberHandler(handlers)

	return &fiberGroup{
		engine: e.engine.Use(fiberHandlers),
	}
}

func (e *fiberGroup) Get(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)

	return &fiberGroup{
		engine: e.engine.Get(path, fiberHandlers...),
	}
}

func (e *fiberGroup) Post(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)

	return &fiberGroup{
		engine: e.engine.Post(path, fiberHandlers...),
	}
}

func (e *fiberGroup) Put(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)

	return &fiberGroup{
		engine: e.engine.Put(path, fiberHandlers...),
	}
}

func (e *fiberGroup) Patch(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)

	return &fiberGroup{
		engine: e.engine.Patch(path, fiberHandlers...),
	}
}

func (e *fiberGroup) Delete(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := toFiberHandler(handlers)

	return &fiberGroup{
		engine: e.engine.Delete(path, fiberHandlers...),
	}
}

// TODO: Restructure the struct and remove this function.
func (e *fiberGroup) Listener(port string) error {
	return nil
}
