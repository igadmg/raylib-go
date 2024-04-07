package main

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

/*******************************************************************************************
*
*   raygui - Controls test
*
*   TEST CONTROLS:
*       - gui.ScrollPanel()
*
*   DEPENDENCIES:
*       raylib 4.0  - Windowing/input management and drawing.
*       raygui 3.0  - Immediate-mode GUI controls.
*
*   COMPILATION (Windows - MinGW):
*       gcc -o $(NAME_PART).exe $(FILE_NAME) -I../../src -lraylib -lopengl32 -lgdi32 -std=c99
*
*   COMPILATION (Linux - gcc):
*       gcc -o $(NAME_PART) $(FILE_NAME) -I../../src -lraylib -lGL -lm -lpthread -ldl -lrt -lX11 -std=c99
*
*   LICENSE: zlib/libpng
*
*   Copyright (c) 2019-2022 Vlad Adrian (@Demizdor) and Ramon Santamaria (@raysan5)
*
**********************************************************************************************/

// ------------------------------------------------------------------------------------
// Program main entry point
// ------------------------------------------------------------------------------------
func main() {

	// Initialization
	//---------------------------------------------------------------------------------------
	const (
		screenWidth  = 800
		screenHeight = 450
	)

	rl.InitWindow(screenWidth, screenHeight, "raygui - gui.ScrollPanel()")

	var (
		panelRec        = rl.NewRectangle(20, 40, 200, 150)
		panelContentRec = rl.NewRectangle(0, 0, 340, 340)
		panelView       = rl.NewRectangle(0, 0, 0, 0)
		panelScroll     = rl.NewVector2(99, -20)
		mouseCell       = rl.NewVector2(0, 0)

		showContentArea = true
	)

	rl.SetTargetFPS(60)
	//---------------------------------------------------------------------------------------

	// Main game loop
	for !rl.WindowShouldClose() {
		// Detect window close button or ESC key

		// Update
		//----------------------------------------------------------------------------------
		// TODO: Implement required update logic
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("[%.1f, %.1f]", panelScroll.X, panelScroll.Y), 4, 4, 20, rl.Red)

		gui.ScrollPanel(panelRec, "", panelContentRec, &panelScroll, &panelView)

		rl.BeginScissorMode(int32(panelView.XY.X), int32(panelView.XY.Y), int32(panelView.WH.X), int32(panelView.WH.Y))
		gui.Grid(rl.NewRectangle(
			float32(panelRec.XY.X+panelScroll.X),
			float32(panelRec.XY.Y+panelScroll.Y),
			float32(panelContentRec.WH.X),
			float32(panelContentRec.WH.Y),
		), "", 16, 3, &mouseCell)
		rl.EndScissorMode()

		if showContentArea {
			rl.DrawRectangle(
				int32(panelRec.XY.X+panelScroll.X),
				int32(panelRec.XY.Y+panelScroll.Y),
				int32(panelContentRec.WH.X),
				int32(panelContentRec.WH.Y),
				rl.Fade(rl.Red, 0.1),
			)
		}

		DrawStyleEditControls()

		showContentArea = gui.CheckBox(rl.NewRectangle(565, 80, 20, 20), "SHOW CONTENT AREA", showContentArea)

		panelContentRec.WH.X = gui.SliderBar(rl.NewRectangle(590, 385, 145, 15),
			"WIDTH",
			fmt.Sprintf("%.1f", panelContentRec.WH.X),
			panelContentRec.WH.X,
			1, 600)
		panelContentRec.WH.Y = gui.SliderBar(rl.NewRectangle(590, 410, 145, 15),
			"HEIGHT",
			fmt.Sprintf("%.1f", panelContentRec.WH.Y),
			panelContentRec.WH.Y, 1, 400)

		rl.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	rl.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}

// Draw and process scroll bar style edition controls
func DrawStyleEditControls() {
	// ScrollPanel style controls
	//----------------------------------------------------------
	gui.GroupBox(rl.NewRectangle(550, 170, 220, 205), "SCROLLBAR STYLE")

	var style int32

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.BORDER_WIDTH))
	gui.Label(rl.NewRectangle(555, 195, 110, 10), "BORDER_WIDTH")
	gui.Spinner(rl.NewRectangle(670, 190, 90, 20), "", &style, 0, 6, false)
	gui.SetStyle(gui.SCROLLBAR, gui.BORDER_WIDTH, style)

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.ARROWS_SIZE))
	gui.Label(rl.NewRectangle(555, 220, 110, 10), "ARROWS_SIZE")
	gui.Spinner(rl.NewRectangle(670, 215, 90, 20), "", &style, 4, 14, false)
	gui.SetStyle(gui.SCROLLBAR, gui.ARROWS_SIZE, style)

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.SLIDER_PADDING))
	gui.Label(rl.NewRectangle(555, 245, 110, 10), "SLIDER_PADDING")
	gui.Spinner(rl.NewRectangle(670, 240, 90, 20), "", &style, 0, 14, false)
	gui.SetStyle(gui.SCROLLBAR, gui.SLIDER_PADDING, style)

	style = boolToint32(gui.CheckBox(rl.NewRectangle(565, 280, 20, 20), "ARROWS_VISIBLE", int32Tobool(int32(gui.GetStyle(gui.SCROLLBAR, gui.ARROWS_VISIBLE)))))
	gui.SetStyle(gui.SCROLLBAR, gui.ARROWS_VISIBLE, style)

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.SLIDER_PADDING))
	gui.Label(rl.NewRectangle(555, 325, 110, 10), "SLIDER_PADDING")
	gui.Spinner(rl.NewRectangle(670, 320, 90, 20), "", &style, 0, 14, false)
	gui.SetStyle(gui.SCROLLBAR, gui.SLIDER_PADDING, style)

	style = int32(gui.GetStyle(gui.SCROLLBAR, gui.SLIDER_WIDTH))
	gui.Label(rl.NewRectangle(555, 350, 110, 10), "SLIDER_WIDTH")
	gui.Spinner(rl.NewRectangle(670, 345, 90, 20), "", &style, 2, 100, false)
	gui.SetStyle(gui.SCROLLBAR, gui.SLIDER_WIDTH, style)

	var text string
	if gui.GetStyle(gui.LISTVIEW, gui.SCROLLBAR_SIDE) == gui.SCROLLBAR_LEFT_SIDE {
		text = "SCROLLBAR: LEFT"
	} else {
		text = "SCROLLBAR: RIGHT"
	}
	style = boolToint32(gui.Toggle(rl.NewRectangle(560, 110, 200, 35), text, int32Tobool(int32(gui.GetStyle(gui.LISTVIEW, gui.SCROLLBAR_SIDE)))))
	gui.SetStyle(gui.LISTVIEW, gui.SCROLLBAR_SIDE, style)
	//----------------------------------------------------------

	// ScrollBar style controls
	//----------------------------------------------------------
	gui.GroupBox(rl.NewRectangle(550, 20, 220, 135), "SCROLLPANEL STYLE")

	style = int32(gui.GetStyle(gui.LISTVIEW, gui.SCROLLBAR_WIDTH))
	gui.Label(rl.NewRectangle(555, 35, 110, 10), "SCROLLBAR_WIDTH")
	gui.Spinner(rl.NewRectangle(670, 30, 90, 20), "", &style, 6, 30, false)
	gui.SetStyle(gui.LISTVIEW, gui.SCROLLBAR_WIDTH, style)

	style = int32(gui.GetStyle(gui.DEFAULT, gui.BORDER_WIDTH))
	gui.Label(rl.NewRectangle(555, 60, 110, 10), "BORDER_WIDTH")
	gui.Spinner(rl.NewRectangle(670, 55, 90, 20), "", &style, 0, 20, false)
	gui.SetStyle(gui.DEFAULT, gui.BORDER_WIDTH, style)
	//----------------------------------------------------------
}

func boolToint32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func int32Tobool(v int32) bool {
	return 0 < v
}
