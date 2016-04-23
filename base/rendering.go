package base

import (
	"bytes"

	"github.com/labstack/echo"
	"github.com/lcaballero/gel"
)

const DefaultDocType = "<!DOCTYPE html>\n"

type RenderHtml struct {
	DocType string
	Indent  gel.Indent
}

func (r RenderHtml) ActiveDocType() string {
	if r.DocType == "" {
		return DefaultDocType
	}
	return r.DocType
}

func (r RenderHtml) ToHtml(c echo.Context, httpCode int, view gel.Viewable) error {
	buf := bytes.NewBuffer([]byte{})
	if r.Indent.HasIndent() {
		view.ToView().ToNode().WriteToIndented(r.Indent, buf)
	} else {
		view.ToView().ToNode().WriteTo(buf)
	}

	return c.HTML(httpCode, buf.String())
}
