package app

import (
	"github.com/gofiber/fiber/v3"
)

type Getter interface {
	Value(key any) any
}

type globalItem struct {
	key   any
	value any
}
type AppGlobals struct {
	globals []globalItem
}

func (app *AppGlobals) Set(key, value any) {
	app.globals = append(app.globals, globalItem{key, value})
}
func (app *AppGlobals) Value(key any) any {
	for i := len(app.globals) - 1; i >= 0; i-- {
		if key == app.globals[i].key {
			return app.globals[i].value
		}
	}
	return nil
}

var _ Getter = &AppGlobals{}

func NewGlobals() *AppGlobals {
	return &AppGlobals{
		globals: nil,
	}
}

type AppGlobalsKeyType int

const (
	AppGlobalsKey AppGlobalsKeyType = 0xFF00AA
)

func FromContext(ctx fiber.Ctx) *AppGlobals {
	val := ctx.Locals(AppGlobalsKey)
	if val == nil {
		return nil
	}
	return val.(*AppGlobals)
}

func Value[T any](ctx Getter, key any) T {
	var empty T
	if ctx == nil {
		return empty
	}
	val := ctx.Value(key)
	if casted, ok := val.(T); ok {
		return casted
	}
	return empty
}
