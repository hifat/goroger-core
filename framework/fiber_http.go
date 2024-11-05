package framework

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

func (h *httpFiberCtx) BodyParser(v interface{}) error {
	return h.ctx.BodyParser(v)
}

func (h *httpFiberCtx) JSON(status int, v interface{}) {
	if err := h.ctx.Status(status).JSON(v); err != nil {
		log.Fatal(err)
	}
}

func (h *httpFiberCtx) Abort(status int, message string) {
	h.ctx.Status(status).SendString(message)
}

func (h *httpFiberCtx) AbortWithJSON(status int, v interface{}) {
	if err := h.ctx.Status(status).JSON(v); err != nil {
		log.Fatal(err)
	}
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

func (h *httpFiberCtx) QueryParser(v interface{}) error {
	return h.ctx.QueryParser(v)
}

func (h *httpFiberCtx) Next() {
	if err := h.ctx.Next(); err != nil {
		log.Fatal(err)
	}
}
