package main

import (
	"fmt"
	"image/color"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MENU      = 30
	BORDERS   = 10
	FIELDSIZE = 40
)

func main() {
	game := NewGame()
	screenWidth := int32((game.FieldSize.X)*FIELDSIZE) + 2*BORDERS
	screenHeight := int32((game.FieldSize.Y)*FIELDSIZE) + 2*BORDERS + MENU

	rl.InitWindow(screenWidth, screenHeight, "Snake")

	rl.SetTargetFPS(60)

	t := time.Now()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		c := game.InputRead()

		if time.Now().Sub(t) >= 200*time.Millisecond || c {
			game.GameStep()
			t = time.Now()
		}
		rl.ClearBackground(rl.White)
		game.DrawField()
		game.DrawCheckpoint()
		game.DrawSnake()
		game.DrawMenu()

		rl.EndDrawing()

	}

	rl.CloseWindow()

}

// For android
func init() {
	rl.SetCallbackFunc(main)
}

func (g *Game) InputRead() bool {
	clicked := false
	if g.Running {
		if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) {
			clicked = g.Snake.ChangeDirection(LEFT)
			return clicked
		}
		if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) {
			clicked = g.Snake.ChangeDirection(RIGHT)
			return clicked

		}
		if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
			clicked = g.Snake.ChangeDirection(UP)
			return clicked

		}
		if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) {
			clicked = g.Snake.ChangeDirection(DOWN)
			return clicked
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) && g.GameOver {
		g.NewSnake()
		g.GameOver = false
		g.StartTime = time.Now()
		g.Running = true

		return clicked
	}

	if rl.IsKeyPressed(rl.KeySpace) && !g.GameOver {
		g.Running = !g.Running
		return clicked
	}

	return clicked
}

// Draw - draw game
func (g *Game) DrawField() {

	for i := 0; i < g.FieldSize.X; i++ {
		for j := 0; j < g.FieldSize.Y; j++ {
			if (i+j)%2 == 0 {
				rl.DrawRectangle(RectX(i), RectY(j), FIELDSIZE, FIELDSIZE, rl.Gray)
				rl.DrawRectangle(RectX(i)+1, RectY(j)+1, FIELDSIZE-2, FIELDSIZE-2, rl.White)
			} else {
				rl.DrawRectangle(RectX(i), RectY(j), FIELDSIZE, FIELDSIZE, rl.DarkGray)
				rl.DrawRectangle(RectX(i)+1, RectY(j)+1, FIELDSIZE-2, FIELDSIZE-2, rl.White)
			}
		}
	}

}

func (g *Game) DrawSnake() {

	var eyex1, eyey1, eyex2, eyey2 int32
	var grad float32

	var c1, c2 color.RGBA

	switch g.Snake.Direction {
	case UP:
		eyex1 = int32(0.4 * FIELDSIZE / 2)
		eyey1 = int32(-0.4 * FIELDSIZE / 2)
		eyex2 = -int32(0.4 * FIELDSIZE / 2)
		eyey2 = int32(-0.4 * FIELDSIZE / 2)
		grad = 0

	case DOWN:
		eyex1 = int32(0.4 * FIELDSIZE / 2)
		eyey1 = int32(0.4 * FIELDSIZE / 2)
		eyex2 = -int32(0.4 * FIELDSIZE / 2)
		eyey2 = int32(0.4 * FIELDSIZE / 2)
		grad = 180
	case LEFT:
		eyex1 = -int32(0.4 * FIELDSIZE / 2)
		eyey1 = int32(0.4 * FIELDSIZE / 2)
		eyex2 = -int32(0.4 * FIELDSIZE / 2)
		eyey2 = int32(-0.4 * FIELDSIZE / 2)
		grad = 90

	case RIGHT:
		eyex1 = int32(0.4 * FIELDSIZE / 2)
		eyey1 = int32(0.4 * FIELDSIZE / 2)
		eyex2 = int32(0.4 * FIELDSIZE / 2)
		eyey2 = int32(-0.4 * FIELDSIZE / 2)
		grad = 270

	}

	c1 = GetColor(len(g.Snake.Body) / 10)
	c2 = GetColor(len(g.Snake.Body)/10 + 1 + len(g.Snake.Body)/50 + len(g.Snake.Body)/100)

	for k, v := range g.Snake.Body {

		if k == 0 {
			rl.DrawCircleSector(rl.Vector2{X: float32(RectXCenter(v.X)), Y: float32(RectYCenter(v.Y))}, FIELDSIZE, 45+grad, 135+grad, 3, c1)
			rl.DrawCircleSector(rl.Vector2{X: float32(RectXCenter(v.X)), Y: float32(RectYCenter(v.Y))}, FIELDSIZE, 225+grad, 315+grad, 2, c1)
		}

		if k == 0 {

			rl.DrawCircle(RectXCenter(v.X), RectYCenter(v.Y), FIELDSIZE/1.7, c1)

			rl.DrawCircle(RectXCenter(v.X)+eyex1, RectYCenter(v.Y)+eyey1, FIELDSIZE/3+2, rl.Black)
			rl.DrawCircle(RectXCenter(v.X)+eyex2, RectYCenter(v.Y)+eyey2, FIELDSIZE/3+2, rl.Black)

			rl.DrawCircle(RectXCenter(v.X)+eyex1, RectYCenter(v.Y)+eyey1, FIELDSIZE/3, rl.White)
			rl.DrawCircle(RectXCenter(v.X)+eyex2, RectYCenter(v.Y)+eyey2, FIELDSIZE/3, rl.White)

			rl.DrawCircle(RectXCenter(v.X)+eyex1, RectYCenter(v.Y)+eyey1, FIELDSIZE/6, rl.Black)
			rl.DrawCircle(RectXCenter(v.X)+eyex2, RectYCenter(v.Y)+eyey2, FIELDSIZE/6, rl.Black)

		} else if k%2 == 0 {
			rl.DrawCircle(RectXCenter(v.X), RectYCenter(v.Y), FIELDSIZE/2, c1)
		} else {
			rl.DrawCircle(RectXCenter(v.X), RectYCenter(v.Y), FIELDSIZE/2, c2)
		}

	}

}

func (g *Game) DrawCheckpoint() {
	c1 := GetColor(len(g.Snake.Body)/10 + 1 + len(g.Snake.Body)/50 + len(g.Snake.Body)/100 + 1)

	if !g.Checkpoint.Taken {
		rl.DrawCircle(RectXCenter(g.Checkpoint.Point.X), RectYCenter(g.Checkpoint.Point.Y), FIELDSIZE/2, c1)
	}
}

func (g *Game) DrawMenu() {

	text := fmt.Sprintf("Snake: %v time: %v s", g.SnakeLength(), int(time.Now().Sub(g.StartTime).Seconds()))
	rl.DrawText(text, int32(BORDERS), 5, int32(FIELDSIZE), rl.Black)

	if g.GameOver {
		text := fmt.Sprintf("        GAME OVER \n press SPACE for new game ")

		rl.DrawRectangle(int32(BORDERS), int32(g.FieldSize.Y*FIELDSIZE)/2+int32(BORDERS)+int32(MENU)-16, int32(g.FieldSize.X*FIELDSIZE), 120, rl.Red)
		rl.DrawText(text, int32(g.FieldSize.X*FIELDSIZE)/2+int32(BORDERS)-4*60, int32(g.FieldSize.Y*FIELDSIZE)/2+int32(MENU)+int32(BORDERS), 40, rl.Black)
	}

	if !g.Running {
		text := fmt.Sprintf("Start Game \n press SPACE ")

		rl.DrawRectangle(int32(BORDERS), int32(g.FieldSize.Y*FIELDSIZE)/2+int32(BORDERS)+int32(MENU)-16, int32(g.FieldSize.X*FIELDSIZE), 120, rl.Green)
		rl.DrawText(text, int32(g.FieldSize.X*FIELDSIZE)/2+int32(BORDERS)-4*60, int32(g.FieldSize.Y*FIELDSIZE)/2+int32(MENU)+int32(BORDERS), 40, rl.Black)
	}

}

func RectX(i int) int32 {
	return int32(i*FIELDSIZE) + int32(BORDERS)
}

func RectY(i int) int32 {
	return int32(i*FIELDSIZE) + int32(BORDERS) + int32(MENU)
}

func RectXCenter(i int) int32 {
	return int32(i*FIELDSIZE+FIELDSIZE/2) + int32(BORDERS)
}

func RectYCenter(i int) int32 {
	return int32(i*FIELDSIZE+FIELDSIZE/2) + int32(BORDERS) + int32(MENU)
}

func GetColor(i int) color.RGBA {
	c := []color.RGBA{
		rl.Gray,
		rl.DarkGray,
		rl.Beige,
		rl.DarkBlue,
		rl.DarkBrown,
		rl.DarkGreen,
		rl.DarkPurple,
		rl.Brown,
		rl.Gold,
		rl.Magenta,
		rl.Maroon,
	}

	return c[i%10]
}
