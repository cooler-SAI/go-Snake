package controls

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

func HandleInput() string {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp):
		log.Info().Msg("Moving UP")
		return "UP"
	case ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown):
		log.Info().Msg("Moving DOWN")
		return "DOWN"
	case ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft):
		log.Info().Msg("Moving LEFT")
		return "LEFT"
	case ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight):
		log.Info().Msg("Moving RIGHT")
		return "RIGHT"
	default:
		return ""
	}
}
