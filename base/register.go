package base

import (
	"github.com/labstack/echo"
)

type Register struct {
	Echo *echo.Echo
}

func NewRegister() *Register {
	return &Register{
		Echo: echo.New(),
	}
}
func (r Register) Connect(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Connect(path, handler, m...)
}
func (r Register) Delete(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Delete(path, handler, m...)
}
func (r Register) Get(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Get(path, handler, m...)
}
func (r Register) Head(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Head(path, handler, m...)
}
func (r Register) Options(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Options(path, handler, m...)
}
func (r Register) Patch(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Patch(path, handler, m...)
}
func (r Register) Post(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Post(path, handler, m...)
}
func (r Register) Put(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Put(path, handler, m...)
}
func (r Register) Trace(path string, handler echo.HandlerFunc, m ...echo.Middleware) {
	r.Echo.Trace(path, handler, m...)
}
