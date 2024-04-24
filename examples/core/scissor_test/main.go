package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenW = int32(800)
	screenH = int32(450)
)

func main() {

	rl.InitWindow(screenW, screenH, "raylib [core] example - scissor test")

	scissorArea := rl.NewRectangle(0, 0, 300, 300)
	scissorMode := true

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyS) {
			scissorMode = !scissorMode
		}

		scissorArea.XY.X = float32(rl.GetMouseX())
		scissorArea.XY.Y = float32(rl.GetMouseY())

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if scissorMode {
			rl.BeginScissorMode(scissorArea.ToInt32().XY.X, scissorArea.ToInt32().XY.Y, scissorArea.ToInt32().WH.X, scissorArea.ToInt32().WH.Y)
		}

		rl.DrawRectangle(0, 0, screenW, screenH, rl.Red)
		rl.DrawText("MOVE MOUSE TO REVEAL TEXT", 190, 200, 20, rl.Black)

		if scissorMode {
			rl.EndScissorMode()
		}

		rl.DrawRectangleLinesEx(scissorArea, 1, rl.Black)
		rl.DrawText("S KEY TO TOGGLE MODE", 10, 10, 20, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
