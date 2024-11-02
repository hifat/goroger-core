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

func (h *httpFiberCtx) JSON(status int, v interface{}) error {
	return h.ctx.Status(status).JSON(v)
}

func (h *httpFiberCtx) Abort(status int, message string) {
	h.ctx.Status(status).SendString(message)
}

func (h *httpFiberCtx) SendString(message string) {
	h.ctx.SendString(message)
}

func (h *httpFiberCtx) Set(key string, value interface{}) {
	h.ctx.Locals(key, value)
}

func (h *httpFiberCtx) Get(key string) interface{} {
	return h.ctx.Locals(key)
}

func (h *httpFiberCtx) MustGet(key string) (interface{}, bool) {
	return h.ctx.Locals(key), true
}
