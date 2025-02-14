package main

import (
	"go-Snake/engine"
	"go-Snake/logger"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	logger.Init()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")

	game := &engine.Game{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
