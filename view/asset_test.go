package view

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHashedAssetPath(t *testing.T) {
	Convey("Should resolve and hash asset", t, func() {
		var resolver PathResolver = func(p string) string {
			return ".js/" + p
		}
		hashed, err := resolver.HashedAsset("hashed-path.js")
		So(err, ShouldBeNil)
		So(hashed, ShouldNotBeNil)
		So(hashed.Hash, ShouldNotBeEmpty)
		So(hashed.Path, ShouldNotBeEmpty)
		So(hashed.Resolved, ShouldNotBeEmpty)
	})
}
