package main

import (
	"go-Snake/engine"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Go Snake")

	game := engine.NewGame(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
