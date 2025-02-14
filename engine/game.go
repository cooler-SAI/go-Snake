package engine

import (
	"go-Snake/controls"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	GridSize  = 20
	MoveSpeed = 150 * time.Millisecond
)

func NewGame(width, height int) *Game {
	rand.Seed(time.Now().UnixNano())

	snake := &Snake{
		Dir:  DirectionRight,
		Body: make([][2]int, 0),
	}

	// Initial position
	startX := width / 2 / GridSize * GridSize
	startY := height / 2 / GridSize * GridSize
	for i := 0; i < 3; i++ {
		snake.Body = append(snake.Body, [2]int{startX - i*GridSize, startY})
	}

	return &Game{
		ScreenWidth:  width,
		ScreenHeight: height,
		Snake:        snake,
		Food:         generateFood(width, height),
		GameOver:     false,
		MoveDelay:    MoveSpeed,
		LastMove:     time.Now(),
	}
}

func generateFood(width, height int) *Food {
	return &Food{
		X: (rand.Intn(width/GridSize-2) + 1) * GridSize,
		Y: (rand.Intn(height/GridSize-2) + 1) * GridSize,
	}
}

func (g *Game) Update() error {
	if g.GameOver {
		if controls.HandleInput() == "RESTART" {
			*g = *NewGame(g.ScreenWidth, g.ScreenHeight)
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
		}
	case "DOWN":
		if g.Snake.Dir != DirectionUp {
			g.Snake.Dir = DirectionDown
		}
	case "LEFT":
		if g.Snake.Dir != DirectionRight {
			g.Snake.Dir = DirectionLeft
		}
	case "RIGHT":
		if g.Snake.Dir != DirectionLeft {
			g.Snake.Dir = DirectionRight
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

	// Wall collision
	if head[0] < 0 || head[0] >= g.ScreenWidth ||
		head[1] < 0 || head[1] >= g.ScreenHeight {
		g.GameOver = true
		return
	}

	// Self collision
	for _, segment := range g.Snake.Body[1:] {
		if head[0] == segment[0] && head[1] == segment[1] {
			g.GameOver = true
			return
		}
	}
}

func (g *Game) checkFood() {
	head := g.Snake.Body[0]
	if head[0] == g.Food.X && head[1] == g.Food.Y {
		// Grow snake
		g.Snake.Body = append(g.Snake.Body, g.Snake.Body[len(g.Snake.Body)-1])
		g.Food = generateFood(g.ScreenWidth, g.ScreenHeight)
		g.Score++
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw snake
	for _, segment := range g.Snake.Body {
		ebitenutil.DrawRect(
			screen,
			float64(segment[0]),
			float64(segment[1]),
			GridSize-1,
			GridSize-1,
			SnakeColor,
		)
	}

	// Draw food
	ebitenutil.DrawRect(
		screen,
		float64(g.Food.X),
		float64(g.Food.Y),
		GridSize-1,
		GridSize-1,
		FoodColor,
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
