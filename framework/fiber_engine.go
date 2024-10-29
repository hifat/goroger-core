package framework

import (
	"github.com/gofiber/fiber/v2"
	core "github.com/hifat/goroger-core"
)

type httpEngine struct {
	engine *fiber.App
}

func NewHttpFiberEngine(engine *fiber.App) core.IHttpEngine {
	return &httpEngine{engine}
}

func (e *httpEngine) Get(path string, handlers ...func(core.IHttpCtx)) core.IHttpRouter {
	fiberHandlers := make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		fiberHandlers[i] = func(ctx *fiber.Ctx) error {
			handler(NewHttpFiber(ctx))
			return nil
		}
	}

	return e.engine.Get(path, fiberHandlers...)
}

func (e *httpEngine) Listener(port string) error {
	return e.engine.Listen(port)
}
