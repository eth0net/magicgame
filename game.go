package main

import (
	"github.com/EngoEngine/engo"
	"github.com/raziel2244/magicgame/scenes/game"
)

func main() {
	opts := engo.RunOptions{
		Title:                      "Magic Game",
		Width:                      800,
		Height:                     600,
		GlobalScale:                engo.Point{X: 2, Y: 2},
		ApplicationMajorVersion:    0,
		ApplicationMinorVersion:    3,
		ApplicationRevisionVersion: 0,
	}

	engo.Run(opts, &game.Scene{})
}
