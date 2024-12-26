package framework

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	core "github.com/hifat/goroger-core"
	"github.com/stretchr/testify/assert"
)

func TestNewFiberEngineCtx(t *testing.T) {
	app := fiber.New()
	engine := NewFiberEngineCtx(app)
	assert.NotNil(t, engine)
}

func TestHttpEngine_Use(t *testing.T) {
	app := fiber.New()
	engine := NewFiberEngineCtx(app)

	middlewareCalled := false
	middleware := func(ctx core.IHttpCtx) {
		middlewareCalled = true
	}

	engine.Use(middleware)

	req := httptest.NewRequest("GET", "/", nil)
	_, err := app.Test(req)

	assert.NoError(t, err)
	assert.True(t, middlewareCalled)
}

func TestHttpEngine_Routes(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		register func(engine core.IHttpRouter, path string, handler func(core.IHttpCtx))
	}{
		{"GET", "GET", "/test", func(e core.IHttpRouter, p string, h func(core.IHttpCtx)) { e.Get(p, h) }},
		{"POST", "POST", "/test", func(e core.IHttpRouter, p string, h func(core.IHttpCtx)) { e.Post(p, h) }},
		{"PUT", "PUT", "/test", func(e core.IHttpRouter, p string, h func(core.IHttpCtx)) { e.Put(p, h) }},
		{"PATCH", "PATCH", "/test", func(e core.IHttpRouter, p string, h func(core.IHttpCtx)) { e.Patch(p, h) }},
		{"DELETE", "DELETE", "/test", func(e core.IHttpRouter, p string, h func(core.IHttpCtx)) { e.Delete(p, h) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			engine := NewFiberEngineCtx(app)

			handlerCalled := false
			handler := func(ctx core.IHttpCtx) {
				handlerCalled = true
			}

			tt.register(engine, tt.path, handler)

			req := httptest.NewRequest(tt.method, tt.path, nil)
			_, err := app.Test(req)

			assert.NoError(t, err)
			assert.True(t, handlerCalled)
		})
	}
}

func TestHttpEngine_Listener(t *testing.T) {
	app := fiber.New()
	engine := NewFiberEngineCtx(app)

	// Start the server in a goroutine
	go func() {
		err := engine.Listener(":0")
		assert.NoError(t, err)
	}()

	// Give the server time to start
	app.ShutdownWithTimeout(1)
}
