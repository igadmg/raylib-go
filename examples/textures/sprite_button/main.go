/*******************************************************************************************
*
*   raylib [textures] example - sprite button
*
*   Example originally created with raylib 2.5, last time updated with raylib 2.5
*
*   Example licensed under an unmodified zlib/libpng license, which is an OSI-certified,
*   BSD-like license that allows static linking with closed source software
*
*   Copyright (c) 2019-2024 Ramon Santamaria (@raysan5)
*
********************************************************************************************/
package main

import (
	"github.com/igadmg/gamemath/rect2"
	"github.com/igadmg/gamemath/vector2"
	rl "github.com/igadmg/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
	numFrames    = 3 // Number of frames (rectangles) for the button sprite texture
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [textures] example - sprite button")

	rl.InitAudioDevice() // Initialize audio device

	fxButton := rl.LoadSound("buttonfx.wav") // Load button sound
	button := rl.LoadTexture("button.png")   // Load button texture

	// Define frame rectangle for drawing
	frameHeight := float32(button.Height) / numFrames
	sourceRec := rect2.NewFloat32(vector2.Zero[float32](), vector2.NewFloat32(button.Width, frameHeight))

	// Define button bounds on screen
	btnBounds := rect2.NewFloat32(
		vector2.NewFloat32(screenWidth/2.0-button.Width/2.0, screenHeight/2.0-button.Height/numFrames/2.0),
		vector2.NewFloat32(button.Width, frameHeight),
	)

	btnState := 0      // Button state: 0-NORMAL, 1-MOUSE_HOVER, 2-PRESSED
	btnAction := false // Button action should be activated

	rl.SetTargetFPS(60)

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		mousePoint := rl.GetMousePosition()
		btnAction = false

		// Check button state
		if rl.CheckCollisionPointRec(mousePoint, btnBounds) {
			if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				btnState = 2
			} else {
				btnState = 1
			}

			if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
				btnAction = true
			}
		} else {

			btnState = 0
		}

		if btnAction {
			rl.PlaySound(fxButton)

			// TODO: Any desired action
		}

		// Calculate button frame rectangle to draw depending on button state
		sourceRec.SetY(float32(btnState) * frameHeight)

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTextureRec(button, sourceRec, btnBounds.Position, rl.White) // Draw button frame

		rl.EndDrawing()
	}

	// De-Initialization
	rl.UnloadTexture(&button) // Unload button texture
	rl.UnloadSound(&fxButton) // Unload sound

	rl.CloseAudioDevice() // Close audio device

	rl.CloseWindow() // Close window and OpenGL context
}
