package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth  = int32(1280)
	screenHeight = int32(720)

	pause     = false
	collision = false

	boxCollision     = rl.Rectangle{}
	screenUpperLimit = float32(40)
	boxAspeedX       = float32(4)
)

func main() {

	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - collision area")

	rl.SetTargetFPS(60)

	boxA := rl.NewRectangle(10, float32(rl.GetScreenHeight()/2)-50, 200, 100)
	boxB := rl.NewRectangle(float32(rl.GetScreenWidth()/2)-30, float32(rl.GetScreenHeight()/2)-30, 60, 60)

	for !rl.WindowShouldClose() {

		mousePos := rl.GetMousePosition()

		if !pause {
			boxA.XY.X += boxAspeedX
		}
		if boxA.XY.X+boxA.WH.X >= float32(rl.GetScreenWidth()) || boxA.XY.X <= 0 {
			boxAspeedX *= -1
		}

		boxB.XY.X = mousePos.X - boxB.WH.X/2
		boxB.XY.Y = mousePos.Y - boxB.WH.Y/2

		if boxB.XY.X+boxB.WH.X >= float32(rl.GetScreenWidth()) {
			boxB.XY.X = float32(rl.GetScreenWidth()) - boxB.WH.X
		} else if boxB.XY.X <= 0 {
			boxB.XY.X = 0
		}

		if boxB.XY.Y+boxB.WH.Y >= float32(rl.GetScreenHeight()) {
			boxB.XY.Y = float32(rl.GetScreenHeight()) - boxB.WH.Y
		} else if boxB.XY.X <= screenUpperLimit {
			boxB.XY.Y = screenUpperLimit
		}

		collision := rl.CheckCollisionRecs(boxA, boxB)

		if collision {
			boxCollision = rl.GetCollisionRec(boxA, boxB)
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if collision {
			rl.DrawRectangle(0, 0, screenWidth, int32(screenUpperLimit), rl.Red)
			rl.DrawRectangleRec(boxCollision, rl.Lime)
			rl.DrawText("COLLISION", int32(rl.GetScreenWidth()/2)-(rl.MeasureText("COLLISION", 20)/2), int32(screenUpperLimit/2)-10, 20, rl.Black)
			rl.DrawText("Collision Area: "+fmt.Sprint(boxCollision.WH.X*boxCollision.WH.Y), int32(rl.GetScreenWidth()/2)-100, int32(screenUpperLimit+10), 20, rl.Black)
		} else {
			rl.DrawRectangle(0, 0, screenWidth, int32(screenUpperLimit), rl.Black)
		}

		rl.DrawRectangleRec(boxA, rl.Orange)
		rl.DrawRectangleRec(boxB, rl.Blue)

		rl.DrawText("Press SPACE to PAUSE/RESUME", 20, int32(rl.GetScreenHeight())-35, 20, rl.Black)

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
