package controls

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func HandleInput() string {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW):
		return "UP"
	case ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS):
		return "DOWN"
	case ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA):
		return "LEFT"
	case ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD):
		return "RIGHT"
	case ebiten.IsKeyPressed(ebiten.KeyR):
		return "RESTART"
	}
	return ""
}
