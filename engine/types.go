package engine

type Direction int

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Snake struct {
	X, Y  int       // Текущая позиция головы
	Body  [][2]int  // Сегменты тела
	Dir   Direction // Направление
	Speed int       // Скорость (пикселей в кадр)
}

type Food struct {
	X, Y int // Позиция еды
}

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Snake        *Snake
	Food         *Food
	GameOver     bool
}
