package framework

import (
	"github.com/gofiber/fiber/v2"
	core "github.com/hifat/goroger-core"
)

type httpFiberCtx struct {
	h *fiber.Ctx
}

func NewHttpFiber(h *fiber.Ctx) core.IHttpCtx {
	return &httpFiberCtx{h}
}

func (c *httpFiberCtx) Get(key string) string {
	return c.h.Params(key)
}
