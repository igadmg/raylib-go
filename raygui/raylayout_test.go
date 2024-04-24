package raygui_test

import (
	"testing"

	"github.com/EliCDavis/vector/test"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestCanvasLayout(t *testing.T) {
	cl := rg.CanvasLayout(rl.NewRectangle(-50, -50, 200, 200))
	{
		r := cl.Layout(rl.AnchorTopLeft, rl.AnchorTopLeft, rl.NewRectangle(0, 0, 20, 20))
		test.AssertRectangleInDelta(t, r, rl.NewRectangle(-50, -50, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorTopRight, rl.AnchorTopRight, rl.NewRectangle(0, 0, 20, 20))
		test.AssertRectangleInDelta(t, r, rl.NewRectangle(130, -50, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorTopLeft, rl.AnchorCenter, rl.NewRectangle(0, 0, 20, 20))
		test.AssertRectangleInDelta(t, r, rl.NewRectangle(-60, -60, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorBottomLeft, rl.AnchorBottomLeft, rl.NewRectangle(0, 0, 20, 20))
		test.AssertRectangleInDelta(t, r, rl.NewRectangle(-50, 130, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorBottomRight, rl.AnchorBottomRight, rl.NewRectangle(0, 0, 20, 20))
		test.AssertRectangleInDelta(t, r, rl.NewRectangle(130, 130, 20, 20), 0.00001)
	}
}
