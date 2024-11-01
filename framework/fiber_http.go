package framework

import (
	"github.com/gofiber/fiber/v2"
	core "github.com/hifat/goroger-core"
)

type httpFiberCtx struct {
	ctx *fiber.Ctx
}

func NewHttpFiber(h *fiber.Ctx) core.IHttpCtx {
	return &httpFiberCtx{h}
}

func (h *httpFiberCtx) Param(key string) string {
	return h.ctx.Params(key)
}

func (h *httpFiberCtx) ShouldBind(v interface{}) error {
	return h.ctx.BodyParser(v)
}
