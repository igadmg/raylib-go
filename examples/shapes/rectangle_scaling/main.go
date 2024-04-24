package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	mouseScaleMarkSize              = float32(12)
	rec                             = rl.NewRectangle(100, 100, 200, 80)
	mouseScaleReady, MouseScaleMode = false, false
)

func main() {
	screenWidth := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - rectangle scaling mouse")

	rl.SetTargetFPS(60)

	rl.SetMousePosition(0, 0)

	for !rl.WindowShouldClose() {

		mousePos := rl.GetMousePosition()

		if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(rec.XY.X+rec.WH.X-mouseScaleMarkSize, rec.XY.Y+rec.WH.Y-mouseScaleMarkSize, mouseScaleMarkSize, mouseScaleMarkSize)) {
			mouseScaleReady = true
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				MouseScaleMode = true
			}
		} else {
			mouseScaleReady = false
		}

		if MouseScaleMode {

			mouseScaleReady = true
			rec.WH.X = mousePos.X - rec.XY.X
			rec.WH.Y = mousePos.Y - rec.XY.Y

			// CHECK MIN MAX REC SIZES
			if rec.WH.X < mouseScaleMarkSize {
				rec.WH.X = rec.WH.Y
			}
			if rec.WH.Y < mouseScaleMarkSize {
				rec.WH.Y = rec.WH.X
			}
			if rec.WH.X > (float32(rl.GetScreenWidth()) - rec.XY.X) {
				rec.WH.X = float32(rl.GetScreenWidth()) - rec.XY.X
			}
			if rec.WH.Y > (float32(rl.GetScreenHeight()) - rec.XY.Y) {
				rec.WH.Y = float32(rl.GetScreenHeight()) - rec.XY.Y
			}
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				MouseScaleMode = false
			}

		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Scale rectangle dragging from bottom-right corner!", 10, 10, 20, rl.Black)

		rl.DrawRectangleRec(rec, rl.Fade(rl.Green, 0.5))

		if mouseScaleReady {
			rl.DrawRectangleLinesEx(rec, 1, rl.Red)
			rl.DrawTriangle(rl.NewVector2(rec.XY.X+rec.WH.X-mouseScaleMarkSize, rec.XY.Y+rec.WH.Y), rl.NewVector2(rec.XY.X+rec.WH.X, rec.XY.Y+rec.WH.Y), rl.NewVector2(rec.XY.X+rec.WH.X, rec.XY.Y+rec.WH.Y-mouseScaleMarkSize), rl.Red)
		}

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
