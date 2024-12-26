package framework

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	return fiber.New()
}

type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestFiberCtx_Param(t *testing.T) {
	app := setupTestApp()
	app.Get("/:id", func(c *fiber.Ctx) error {
		ctx := NewFiberCtx(c)
		assert.Equal(t, "123", ctx.Param("id"))
		return nil
	})

	req := httptest.NewRequest("GET", "/123", nil)
	_, err := app.Test(req)
	assert.NoError(t, err)
}

func TestFiberCtx_BodyParser(t *testing.T) {
	app := setupTestApp()
	app.Post("/", func(c *fiber.Ctx) error {
		ctx := NewFiberCtx(c)
		var data TestStruct
		err := ctx.BodyParser(&data)
		assert.NoError(t, err)
		assert.Equal(t, "test", data.Name)
		assert.Equal(t, 25, data.Age)
		return nil
	})

	body := `{"name":"test","age":25}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	_, err := app.Test(req)
	assert.NoError(t, err)
}

func TestFiberCtx_JSON(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		ctx := NewFiberCtx(c)
		ctx.JSON(200, TestStruct{Name: "test", Age: 25})
		return nil
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	body, _ := io.ReadAll(resp.Body)
	var result TestStruct
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.Equal(t, "test", result.Name)
	assert.Equal(t, 25, result.Age)
}

func TestFiberCtx_Locals(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		ctx := NewFiberCtx(c)
		ctx.Set("key", "value")

		val := ctx.Get("key")
		assert.Equal(t, "value", val)

		mustVal, exists := ctx.MustGet("key")
		assert.True(t, exists)
		assert.Equal(t, "value", mustVal)
		return nil
	})

	req := httptest.NewRequest("GET", "/", nil)
	_, err := app.Test(req)
	assert.NoError(t, err)
}

func TestFiberCtx_QueryParser(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		ctx := NewFiberCtx(c)
		var query TestStruct
		err := ctx.QueryParser(&query)
		assert.NoError(t, err)
		assert.Equal(t, "go", query.Name)
		assert.Equal(t, 25, query.Age)
		return nil
	})

	req := httptest.NewRequest("GET", "/?name=test&age=25", nil)
	_, err := app.Test(req)
	assert.NoError(t, err)
}
