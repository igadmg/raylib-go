package main

import (
	//"fmt"

	"github.com/igadmg/gamemath/vector3"
	rl "github.com/igadmg/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - heightmap loading and drawing")

	camera := rl.Camera{}
	camera.Position = vector3.NewFloat32(18.0, 16.0, 18.0)
	camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)
	camera.Fovy = 45.0

	image := rl.LoadImage("heightmap.png")    // Load heightmap image (RAM)
	texture := rl.LoadTextureFromImage(image) // Convert image to texture (VRAM)

	mesh := rl.GenMeshHeightmap(image, vector3.NewFloat32(16, 8, 16)) // Generate heightmap mesh (RAM and VRAM)
	model := rl.LoadModelFromMesh(mesh)                               // Load model from generated mesh

	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, &texture) // Set map diffuse texture

	mapPosition := vector3.NewFloat32(-8.0, 0.0, -8.0) // Set model position

	rl.UnloadImage(&image) // Unload heightmap image from RAM, already uploaded to VRAM

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update

		rl.UpdateCamera(&camera, rl.CameraOrbital) // Update camera with orbital camera mode

		// Draw

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawModel(model, mapPosition, 1.0, rl.Red)

		rl.DrawGrid(20, 1.0)

		rl.EndMode3D()

		rl.DrawTexture(texture, screenWidth-texture.Width-20, 20, rl.White)
		rl.DrawRectangleLines(screenWidth-texture.Width-20, 20, texture.Width, texture.Height, rl.Green)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.UnloadTexture(&texture) // Unload map texture
	rl.UnloadModel(&model)     // Unload map model

	rl.CloseWindow()
}
