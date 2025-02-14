package main

import (
	"go-Snake/engine"
	"go-Snake/logger"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func main() {
	logger.Init()

	logger.Log.Info().
		Int("width", ScreenWidth).
		Int("height", ScreenHeight).
		Msg("Starting application")

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Go Snake")

	game := engine.NewGame(ScreenWidth, ScreenHeight)

	logger.Log.Info().Msg("Game initialized")

	if err := ebiten.RunGame(game); err != nil {
		logger.Log.Fatal().
			Err(err).
			Msg("Failed to run game")
	}

	logger.Log.Info().Msg("Application exited normally")
}
