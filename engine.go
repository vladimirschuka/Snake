package main

import (
	"math/rand"
	"time"
)

const x = 40
const y = 40

const (
	UP    Direction = "UP"
	DOWN  Direction = "DOWN"
	LEFT  Direction = "LEFT"
	RIGHT Direction = "RIGHT"
)

type Direction string

type Point struct {
	X int
	Y int
}

type Snake struct {
	Direction Direction
	Body      []Point
}

type Game struct {
	Snake      *Snake
	Checkpoint *CheckPoint
	FieldSize  Point
	GameOver   bool
	Running    bool
	StartTime  time.Time
}

type CheckPoint struct {
	Point Point
	Taken bool
}

func (s *Snake) NewHead() Point {
	var newHead Point
	switch s.Direction {
	case UP:
		newHead = Point{
			X: s.Body[0].X,
			Y: s.Body[0].Y - 1,
		}
	case DOWN:
		newHead = Point{
			X: s.Body[0].X,
			Y: s.Body[0].Y + 1,
		}
	case LEFT:
		newHead = Point{
			X: s.Body[0].X - 1,
			Y: s.Body[0].Y,
		}
	case RIGHT:
		newHead = Point{
			X: s.Body[0].X + 1,
			Y: s.Body[0].Y,
		}
	}
	return newHead
}

func (s *Snake) Move(n Point, keepTail bool) {
	if keepTail {
		s.Body = append([]Point{n}, s.Body...)
	} else {
		s.Body = append([]Point{n}, s.Body[:len(s.Body)-1]...)
	}
}

func (s *Snake) ChangeDirection(n Direction) bool {
	changed := false
	if s.Direction.Opposite() != n && s.Direction != n {
		s.Direction = n
		changed = true
	}
	return changed
}

func Equal(p1 Point, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func (f *Game) CheckPointCollisions(p Point) bool {
	return Equal(p, f.Checkpoint.Point)
}

func (f *Game) BorderCollisions(p Point) bool {
	return p.X < 0 || p.Y < 0 || p.X >= f.FieldSize.X || p.Y >= f.FieldSize.Y
}

func (f Game) SnakeEatItself(p Point) bool {
	res := false
	for _, v := range f.Snake.Body {

		res = res || Equal(v, p)
	}
	return res
}

func (d Direction) Opposite() Direction {
	var o Direction
	switch d {
	case UP:
		o = DOWN
	case DOWN:
		o = UP
	case LEFT:
		o = RIGHT
	case RIGHT:
		o = LEFT
	}
	return o
}

func (f *Game) GameStep() {

	n := f.Snake.NewHead()

	g := f.BorderCollisions(n) || f.SnakeEatItself(n)
	if g {
		f.GameOver = true
	}

	if !f.GameOver && f.Running {
		y := f.CheckPointCollisions(n)
		f.Snake.Move(n, y)
		if y {
			f.NewCheckPoint()
		}
	}
}

func NewGame() *Game {
	var g *Game

	g = &Game{
		FieldSize: Point{20, 20},
		GameOver:  false,
		Running:   false,
		StartTime: time.Now(),
	}
	g.NewSnake()
	g.NewCheckPoint()

	return g
}

func (f *Game) NewSnake() {

	x := f.FieldSize.X / 2

	y := f.FieldSize.Y / 2

	f.Snake = &Snake{
		Direction: UP,
		Body: []Point{
			{
				X: x,
				Y: y,
			},
			{
				X: x + 1,
				Y: y,
			},
			{
				X: x + 2,
				Y: y,
			},
			{
				X: x + 3,
				Y: y,
			}},
	}
}

func (g *Game) NewCheckPoint() {

	FreePoints := []Point{}

	for i := 0; i < g.FieldSize.X; i++ {
		for j := 0; j < g.FieldSize.X; j++ {
			pfree := true
			for _, v := range g.Snake.Body {
				if v.X == i && v.Y == j {
					pfree = false
				}
			}
			if pfree {
				FreePoints = append(FreePoints, Point{
					X: i,
					Y: j,
				})
			}
		}
	}

	l := int32(len(FreePoints))

	pnum := rand.Int31n(l)

	g.Checkpoint = &CheckPoint{
		Point: FreePoints[pnum],
		Taken: false,
	}
}

func (g Game) SnakeLength() int {
	return len(g.Snake.Body)
}
