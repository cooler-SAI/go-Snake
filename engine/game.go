package engine

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"go-Snake/controls"
	"go-Snake/logger"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	GridSize  = 20
	MoveSpeed = 150 * time.Millisecond
)

func NewGame(width, height int) *Game {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	snake := &Snake{
		Dir:  DirectionRight,
		Body: make([][2]int, 0),
	}

	startX := width / 2 / GridSize * GridSize
	startY := height / 2 / GridSize * GridSize
	for i := 0; i < 3; i++ {
		snake.Body = append(snake.Body, [2]int{startX - i*GridSize, startY})
	}

	game := &Game{
		ScreenWidth:  width,
		ScreenHeight: height,
		Snake:        snake,
		GameOver:     false,
		MoveDelay:    MoveSpeed,
		LastMove:     time.Now(),
		RNG:          rng,
	}

	game.Food = game.generateFood()
	return game
}

func (g *Game) generateFood() *Food {
	maxX := (g.ScreenWidth / GridSize) - 1
	maxY := (g.ScreenHeight / GridSize) - 1

	return &Food{
		X: g.RNG.Intn(maxX) * GridSize,
		Y: g.RNG.Intn(maxY) * GridSize,
	}
}

func (g *Game) Update() error {
	if g.GameOver {
		if controls.HandleInput() == "RESTART" {
			*g = *NewGame(g.ScreenWidth, g.ScreenHeight)
			logger.Log.Info().Msg("Game restarted")
		}
		return nil
	}

	if time.Since(g.LastMove) < g.MoveDelay {
		return nil
	}

	direction := controls.HandleInput()
	g.updateDirection(direction)
	g.moveSnake()
	g.checkCollisions()
	g.checkFood()

	g.LastMove = time.Now()
	return nil
}

func (g *Game) updateDirection(newDir string) {
	switch newDir {
	case "UP":
		if g.Snake.Dir != DirectionDown {
			g.Snake.Dir = DirectionUp
			logger.Log.Debug().Str("direction", "UP").Msg("Direction changed")
		}
	case "DOWN":
		if g.Snake.Dir != DirectionUp {
			g.Snake.Dir = DirectionDown
			logger.Log.Debug().Str("direction", "DOWN").Msg("Direction changed")
		}
	case "LEFT":
		if g.Snake.Dir != DirectionRight {
			g.Snake.Dir = DirectionLeft
			logger.Log.Debug().Str("direction", "LEFT").Msg("Direction changed")
		}
	case "RIGHT":
		if g.Snake.Dir != DirectionLeft {
			g.Snake.Dir = DirectionRight
			logger.Log.Debug().Str("direction", "RIGHT").Msg("Direction changed")
		}
	}
}

func (g *Game) moveSnake() {
	head := g.Snake.Body[0]
	newHead := [2]int{head[0], head[1]}

	switch g.Snake.Dir {
	case DirectionUp:
		newHead[1] -= GridSize
	case DirectionDown:
		newHead[1] += GridSize
	case DirectionLeft:
		newHead[0] -= GridSize
	case DirectionRight:
		newHead[0] += GridSize
	}

	g.Snake.Body = append([][2]int{newHead}, g.Snake.Body...)
	g.Snake.Body = g.Snake.Body[:len(g.Snake.Body)-1]
}

func (g *Game) checkCollisions() {
	head := g.Snake.Body[0]

	if head[0] < 0 || head[0] >= g.ScreenWidth ||
		head[1] < 0 || head[1] >= g.ScreenHeight {
		g.GameOver = true
		logger.Log.Warn().Interface("position", head).Msg("Wall collision")
		return
	}

	for _, segment := range g.Snake.Body[1:] {
		if head[0] == segment[0] && head[1] == segment[1] {
			g.GameOver = true
			logger.Log.Warn().Interface("position", head).Msg("Self collision")
			return
		}
	}
}

func (g *Game) checkFood() {
	head := g.Snake.Body[0]
	if head[0] == g.Food.X && head[1] == g.Food.Y {
		g.Snake.Body = append(g.Snake.Body, g.Snake.Body[len(g.Snake.Body)-1])
		g.Food = g.generateFood()
		g.Score++
		logger.Log.Info().Int("score", g.Score).Msg("Food eaten")
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw snake
	for _, segment := range g.Snake.Body {
		vector.DrawFilledRect(
			screen,
			float32(segment[0]),
			float32(segment[1]),
			GridSize-1,
			GridSize-1,
			color.RGBA{R: 50, G: 205, B: 50, A: 255},
			false,
		)
	}

	// Draw food
	vector.DrawFilledRect(
		screen,
		float32(g.Food.X),
		float32(g.Food.Y),
		GridSize-1,
		GridSize-1,
		color.RGBA{R: 220, G: 20, B: 60, A: 255},
		false,
	)

	// Draw UI
	ebitenutil.DebugPrint(screen,
		"Score: "+strconv.Itoa(g.Score)+
			"\nUse WASD/Arrows to move"+
			"\nPress R to restart",
	)

	if g.GameOver {
		ebitenutil.DebugPrintAt(screen,
			"GAME OVER!",
			g.ScreenWidth/2-50,
			g.ScreenHeight/2-10,
		)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
