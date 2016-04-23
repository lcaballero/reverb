package example

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	. "github.com/lcaballero/gel"
	"github.com/lcaballero/reverb/base"
	"github.com/lcaballero/reverb/view"
)

func Hello(render base.RenderHtml, toAsset view.AssetProvider) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := view.Root{
			Title:  Title(Text("Reverb")),
			Assets: toAsset("reset.css"),
			Main:   Div(H1(Text("Root"))),
		}
		return render.ToHtml(c, http.StatusOK, r)
	}
}

func Run() {
	var resolver view.PathResolver = view.RootResolver("example")
	render := base.RenderHtml{}

	r := base.NewRegister()
	r.Get("/", Hello(render, resolver.ToAsset))
	r.Get("/asset/:kind/:hash/:file", resolver.ToHandler())

	ip := "127.0.0.1:1313"
	fmt.Printf("Starting web server at: %s\n", ip)
	r.Echo.Run(standard.New(ip))
}
