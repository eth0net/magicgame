package assets

import "embed"

//go:embed spritesheets
//go:embed tilemaps/*min.tmx
//go:embed tilesets/*.png
var fs embed.FS

// ReadFile wrapper for embedded assets fs.
func ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(name)
}
