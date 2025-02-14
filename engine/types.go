package engine

import (
	"image/color"
	"time"
)

type Direction int

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Snake struct {
	Body  [][2]int
	Dir   Direction
	Speed int
}

type Food struct {
	X, Y int
}

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Snake        *Snake
	Food         *Food
	GameOver     bool
	Score        int
	LastMove     time.Time
	MoveDelay    time.Duration
}

var (
	SnakeColor = color.RGBA{0, 255, 0, 255}
	FoodColor  = color.RGBA{255, 0, 0, 255}
)
