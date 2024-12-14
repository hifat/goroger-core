package framework

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	core "github.com/hifat/goroger-core"
)

type fiberCtx struct {
	ctx *fiber.Ctx
}

func NewFiberCtx(h *fiber.Ctx) core.IHttpCtx {
	return &fiberCtx{h}
}

func (h *fiberCtx) Param(key string) string {
	return h.ctx.Params(key)
}

func (h *fiberCtx) BodyParser(v interface{}) error {
	return h.ctx.BodyParser(v)
}

func (h *fiberCtx) JSON(status int, v interface{}) {
	if err := h.ctx.Status(status).JSON(v); err != nil {
		log.Fatal(err)
	}
}

func (h *fiberCtx) Abort(status int, message string) {
	if err := h.ctx.Status(status).SendString(message); err != nil {
		log.Fatal(err)
	}
}

func (h *fiberCtx) AbortWithJSON(status int, v interface{}) {
	if err := h.ctx.Status(status).JSON(v); err != nil {
		log.Fatal(err)
	}
}

func (h *fiberCtx) SendString(message string) {
	if err := h.ctx.SendString(message); err != nil {
		log.Fatal(err)
	}
}

func (h *fiberCtx) Set(key string, value interface{}) {
	h.ctx.Locals(key, value)
}

func (h *fiberCtx) Get(key string) interface{} {
	return h.ctx.Locals(key)
}

func (h *fiberCtx) MustGet(key string) (interface{}, bool) {
	return h.ctx.Locals(key), true
}

func (h *fiberCtx) QueryParser(v interface{}) error {
	return h.ctx.QueryParser(v)
}

func (h *fiberCtx) Next() {
	if err := h.ctx.Next(); err != nil {
		log.Fatal(err)
	}
}
