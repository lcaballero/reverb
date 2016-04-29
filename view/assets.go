package view

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/lcaballero/gel"
	"github.com/labstack/echo"
	"net/http"
)

// StaticAsset represents file based assets like js and css.
type StaticAsset string

// AssetResolver takes a path and provides the code.
type AssetResolver func(path string) StaticAsset

// PathResolver translates a given path to asset paths.
type PathResolver func(path ...string) string

// AssetProvider provides a link or script tag based on the extension,
// hashed to a file.
type AssetProvider func(assetPath string) View

type HashedAsset struct {
	Hash, Path, Resolved, Ext, AssetRoot string
	Resolver PathResolver
}

func Js(src string) View {
	return Script.Atts("type", "text/javascript", "src", src)
}

func Css(src string) View {
	return Link.Atts(
		"link", "text/css",
		"rel", "stylesheet",
		"href", src,
	)
}

func (r PathResolver) ToAsset(filename string) View {
	hashed, err := r.HashedAsset(filename)
	if err != nil {
		fmt.Println(err)
		return None()
	}
	return hashed.ToView()
}

func (resolver PathResolver) HashedAsset(filename string) (*HashedAsset, error) {
	ext := filepath.Ext(filename)[1:]
	file := resolver(ext, filename)

	f, err := os.Lstat(file)
	if os.IsNotExist(err) {
		return nil, err
	}

	if f.IsDir() {
		return nil, errors.New("Required file not dir")
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	hash := fmt.Sprintf("%x", md5.Sum(b))

	hashed := &HashedAsset{
		Resolver:  resolver,
		AssetRoot: "/asset",
		Hash:      hash,
		Path:      filename,
		Resolved:  file,
		Ext:       ext,
	}
	return hashed, nil
}

func (resolve PathResolver) ToHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		kind := c.Param("kind")
		file := c.Param("file")
		hash := c.Param("hash")
		asset := resolve(kind, file)
		res := c.Response()
		headers := res.Header()
		headers.Add("Cache-Control", "max-age=31536000")
		headers.Add("ETag", hash)

		match := c.Request().Header().Get("If-None-Match")
		fmt.Printf("match: %s\n", match)

		if match == hash {
			return c.HTML(http.StatusNotModified, "")
		}
		return c.File(asset)
	}
}

func (h *HashedAsset) ToView() View {
	src := fmt.Sprintf("%s/%s/%s/%s", h.AssetRoot, h.Ext, h.Hash, h.Path)
	switch h.Ext {
	case "js":
		return Js(src)
	case "css":
		return Css(src)
	default:
		return None()
	}
}

func RootResolver(root string) PathResolver {
	var resolver PathResolver = func(parts ...string) string {
		path := []string{ root }
		path = append(path, parts...)
		return filepath.Join(path...)
	}
	return resolver;
}

func (p PathResolver) ServeFile(c echo.Context) error {
	file := c.Param("file")
	img := p(file)
	return c.File(img)
}
