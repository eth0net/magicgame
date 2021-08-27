package assets

import "embed"

//go:embed effects spritesheets tilemaps/*min.tmx tilesets/*.png
var fs embed.FS

// ReadFile wrapper for embedded assets fs.
func ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(name)
}
