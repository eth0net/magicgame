package main

import (
	"github.com/EngoEngine/engo"
	"github.com/raziel2244/magicgame/scenes"
)

func main() {
	opts := engo.RunOptions{
		Title:       "Magic Game",
		Width:       800,
		Height:      600,
		GlobalScale: engo.Point{X: 2, Y: 2},
	}

	engo.Run(opts, &scenes.GameScene{})
}
