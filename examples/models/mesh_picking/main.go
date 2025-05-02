/*******************************************************************************************
*
*   raylib [models] example - Mesh picking in 3d mode, ground plane, triangle, mesh
*
*   Example originally created with raylib 1.7, last time updated with raylib 4.0
*
*   Example contributed by Joel Davis (@joeld42) and reviewed by Ramon Santamaria (@raysan5)
*
*   Example licensed under an unmodified zlib/libpng license, which is an OSI-certified,
*   BSD-like license that allows static linking with closed source software
*
*   Copyright (c) 2017-2024 Joel Davis (@joeld42) and Ramon Santamaria (@raysan5)
*
********************************************************************************************/
package main

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/igadmg/gamemath/vector3"
	rl "github.com/igadmg/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
)

// Program main entry point
func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - mesh picking")

	// Define the camera to look into our 3d world
	camera := rl.Camera{
		Position: vector3.Float32{
			X: 20.0,
			Y: 20.0,
			Z: 20.0,
		}, // Camera position
		Target:     vector3.Float32{Y: 8.0}, // Camera looking at point
		Up:         vector3.Float32{Y: 1.6}, // Camera up vector (rotation towards target)
		Fovy:       45.0,                    // Camera field-of-view Y
		Projection: rl.CameraPerspective,    // Camera projection type
	}
	var ray rl.Ray // Picking ray

	tower := rl.LoadModel("turret.obj")             // Load OBJ model
	texture := rl.LoadTexture("turret_diffuse.png") // Load model texture
	materials := unsafe.Slice(tower.Materials, tower.MaterialCount)
	materials[0].GetMap(rl.MapDiffuse).Texture = texture // Set model diffuse texture

	towerPos := vector3.Float32{} // Set model position
	meshes := unsafe.Slice(tower.Meshes, tower.MeshCount)
	towerBBox := rl.GetMeshBoundingBox(meshes[0]) // Get mesh bounding box

	// Ground quad
	g0 := vector3.Float32{
		X: -50.0,
		Z: -50.0,
	}
	g1 := vector3.Float32{
		X: -50.0,
		Z: 50.0,
	}
	g2 := vector3.Float32{
		X: 50.0,
		Z: 50.0,
	}
	g3 := vector3.Float32{
		X: 50.0,
		Z: -50.0,
	}

	// Test triangle
	ta := vector3.Float32{
		X: -25.0,
		Y: 0.5,
	}
	tb := vector3.Float32{
		X: -4.0,
		Y: 2.5,
		Z: 1.0,
	}
	tc := vector3.Float32{
		X: -8.0,
		Y: 6.5,
	}

	bary := vector3.Float32{}

	// Test sphere
	sp := vector3.Float32{
		X: -30.0,
		Y: 5.0,
		Z: 5.0,
	}
	sr := float32(4.0)

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		if rl.IsCursorHidden() {
			rl.UpdateCamera(&camera, rl.CameraFirstPerson) // Update camera
		}
		// Toggle camera controls
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			if rl.IsCursorHidden() {
				rl.EnableCursor()
			} else {
				rl.DisableCursor()
			}
		}

		// Display information about the closest hit
		collision := rl.RayCollision{
			Distance: math.MaxFloat32,
			Hit:      false,
		}
		hitObjectName := "None"
		cursorColor := rl.White

		// Get ray and test against objects
		// See issue : https://github.com/igadmg/raylib-go/issues/457
		//ray = rl.GetScreenToWorldRay(rl.GetMousePosition(), camera)
		ray = rl.GetMouseRay(rl.GetMousePosition(), camera)

		// Check ray collision against ground quad
		groundHitInfo := rl.GetRayCollisionQuad(ray, g0, g1, g2, g3)

		if (groundHitInfo.Hit) && (groundHitInfo.Distance < collision.Distance) {
			collision = groundHitInfo
			cursorColor = rl.Green
			hitObjectName = "Ground"
		}

		// Check ray collision against test triangle
		triHitInfo := rl.GetRayCollisionTriangle(ray, ta, tb, tc)

		if (triHitInfo.Hit) && (triHitInfo.Distance < collision.Distance) {
			collision = triHitInfo
			cursorColor = rl.Purple
			hitObjectName = "Triangle"

			bary = vector3.Float32Barycenter(collision.Point, ta, tb, tc)
		}

		// Check ray collision against test sphere
		sphereHitInfo := rl.GetRayCollisionSphere(ray, sp, sr)

		if (sphereHitInfo.Hit) && (sphereHitInfo.Distance < collision.Distance) {
			collision = sphereHitInfo
			cursorColor = rl.Orange
			hitObjectName = "Sphere"
		}

		// Check ray collision against bounding box first, before trying the full ray-mesh test
		boxHitInfo := rl.GetRayCollisionBox(ray, towerBBox)

		if (boxHitInfo.Hit) && (boxHitInfo.Distance < collision.Distance) {
			collision = boxHitInfo
			cursorColor = rl.Orange
			hitObjectName = "Box"

			// Check ray collision against model meshes
			meshHitInfo := rl.RayCollision{}
			for m := int32(0); m < tower.MeshCount; m++ {
				// NOTE: We consider the model.transform for the collision check, but
				// it can be checked against any transform Matrix, used when checking against same
				// model drawn multiple times with multiple transforms
				meshHitInfo = rl.GetRayCollisionMesh(ray, meshes[m], tower.Transform)
				if meshHitInfo.Hit {
					// Save the closest hit mesh
					if (!collision.Hit) || (collision.Distance > meshHitInfo.Distance) {
						collision = meshHitInfo
					}

					break // Stop once one mesh collision is detected, the colliding mesh is m
				}
			}

			if meshHitInfo.Hit {
				collision = meshHitInfo
				cursorColor = rl.Orange
				hitObjectName = "Mesh"
			}
		}

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)

		// Draw the tower
		// WARNING: If scale is different from 1.0f,
		// not considered by GetRayCollisionModel()
		rl.DrawModel(tower, towerPos, 1.0, rl.White)

		// Draw the test triangle
		rl.DrawLine3D(ta, tb, rl.Purple)
		rl.DrawLine3D(tb, tc, rl.Purple)
		rl.DrawLine3D(tc, ta, rl.Purple)

		// Draw the test sphere
		rl.DrawSphereWires(sp, sr, 8, 8, rl.Purple)

		// Draw the mesh bbox if we hit it
		if boxHitInfo.Hit {
			rl.DrawBoundingBox(towerBBox, rl.Lime)
		}

		// If we hit something, draw the cursor at the hit point
		if collision.Hit {
			rl.DrawCube(collision.Point, 0.3, 0.3, 0.3, cursorColor)
			rl.DrawCubeWires(collision.Point, 0.3, 0.3, 0.3, rl.Red)

			normalEnd := vector3.Float32{}
			normalEnd.X = collision.Point.X + collision.Normal.X
			normalEnd.Y = collision.Point.Y + collision.Normal.Y
			normalEnd.Z = collision.Point.Z + collision.Normal.Z

			rl.DrawLine3D(collision.Point, normalEnd, rl.Red)
		}

		rl.DrawRay(ray, rl.Maroon)
		rl.DrawGrid(10, 10.0)

		rl.EndMode3D()

		// Draw some debug GUI text
		rl.DrawText(fmt.Sprintf("Hit Object: %s", hitObjectName), 10, 50, 10, rl.Black)

		if collision.Hit {
			ypos := int32(70)
			rl.DrawText(fmt.Sprintf("Distance: %3.2f", collision.Distance), 10, ypos, 10, rl.Black)
			rl.DrawText(Vec2Str("Hit Pos : %3.2f %3.2f %3.2f", collision.Point), 10, ypos+15, 10, rl.Black)
			rl.DrawText(Vec2Str("Hit Norm: %3.2f %3.2f %3.2f", collision.Normal), 10, ypos+30, 10, rl.Black)
			if triHitInfo.Hit && hitObjectName == "Triangle" {
				rl.DrawText(Vec2Str("Barycenter: %3.2f %3.2f %3.2f", bary), 10, ypos+45, 10, rl.Black)
			}
		}

		rl.DrawText("Right click mouse to toggle camera controls", 10, 430, 10, rl.Gray)
		rl.DrawText("(c) Turret 3D model by Alberto Cano", screenWidth-200, screenHeight-20, 10, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	// De-Initialization
	rl.UnloadModel(&tower)     // Unload model
	rl.UnloadTexture(&texture) // Unload texture

	rl.CloseWindow() // Close window and OpenGL context
}

func Vec2Str(format string, vec vector3.Float32) string {
	return fmt.Sprintf(format, vec.X, vec.Y, vec.Z)
}
