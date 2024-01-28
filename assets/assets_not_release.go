//go:build !release

package assets

import "os"

func init() {
	assetFS = os.DirFS("./assets")
}
