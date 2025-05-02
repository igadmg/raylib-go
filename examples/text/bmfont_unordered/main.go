package main

import (
	"fmt"

	"github.com/igadmg/gamemath/vector2"
	rl "github.com/igadmg/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - bmfont unordered loading and drawing")

	// NOTE: Using chars outside the [32..127] limits!
	// NOTE: If a character is not found in the font, it just renders a space
	msg := "ASCII extended characters:\n¡¢£¤¥¦§¨©ª«¬®¯°±²³´µ¶·¸¹º»¼½¾¿ÀÁÂÃÄÅÆ\nÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæ\nçèéêëìíîïðñòóôõö÷øùúûüýþÿ"

	// NOTE: Loaded font has an unordered list of characters (chars in the range 32..255)
	font := rl.LoadFont("fonts/pixantiqua.fnt") // BMFont (AngelCode)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Font name:       PixAntiqua", 40, 50, 20, rl.Gray)
		rl.DrawText(fmt.Sprintf("Font base size:           %d", font.BaseSize), 40, 80, 20, rl.Gray)
		rl.DrawText(fmt.Sprintf("Font chars number:     %d", font.GlyphCount), 40, 110, 20, rl.Gray)

		rl.DrawTextEx(font, msg, vector2.NewFloat32(40, 180), float32(font.BaseSize), 0, rl.Maroon)

		rl.EndDrawing()
	}

	rl.UnloadFont(&font) // AngelCode Font unloading

	rl.CloseWindow()
}
