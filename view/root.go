package view

import (
	. "github.com/lcaballero/gel"
)

type Root struct {
	Assets      View
	Scripts     View
	Main        View
	BannerLinks View
	Title       View
}

func (r Root) ToView() View {
	return Html(
		Head(
			Meta.Atts("charset", "UTF-8"),
			Default(r.Title, Title(Text("New Page"))),
			Link.Atts("rel", "shortcut icon", "href", "/favicon.ico", "type", "image/x-icon"),
			Link.Atts("rel", "icon", "href", "/favicon.ico", "type", "image/x-icon"),
			Maybe(r.Assets),
		),
		Body(
			Div.Class("container")(
				Maybe(r.BannerLinks),
				Maybe(r.Main),
			),
			Maybe(r.Scripts),
		),
	)
}
