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
func (r Register) Connect(path string, handler echo.HandlerFunc) {
	r.Echo.Connect(path, handler)
}
func (r Register) Delete(path string, handler echo.HandlerFunc) {
	r.Echo.Delete(path, handler)
}
func (r Register) Get(path string, handler echo.HandlerFunc) {
	r.Echo.Get(path, handler)
}
func (r Register) Head(path string, handler echo.HandlerFunc) {
	r.Echo.Head(path, handler)
}
func (r Register) Options(path string, handler echo.HandlerFunc) {
	r.Echo.Options(path, handler)
}
func (r Register) Patch(path string, handler echo.HandlerFunc) {
	r.Echo.Patch(path, handler)
}
func (r Register) Post(path string, handler echo.HandlerFunc) {
	r.Echo.Post(path, handler)
}
func (r Register) Put(path string, handler echo.HandlerFunc) {
	r.Echo.Put(path, handler)
}
func (r Register) Trace(path string, handler echo.HandlerFunc) {
	r.Echo.Trace(path, handler)
}
