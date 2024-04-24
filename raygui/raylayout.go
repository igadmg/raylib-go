package raygui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Justyfy int

const (
	JustifyTop Justyfy = iota
	JustifyCenter
	JustifyBottom
	JustifyFill
	JustifyLeft
	JustifyRight
)

func (j Justyfy) Justyfy(v, max float32) (nv, dv float32) {
	switch j {
	case JustifyTop:
	case JustifyLeft:
		return v, 0
	case JustifyCenter:
		return v, (max - v) / 2
	case JustifyBottom:
	case JustifyRight:
		return v, max - v
	case JustifyFill:
		return max, 0
	}

	return v, 0
}

type layout struct {
	Bounds rl.Rectangle
}

type canvasLayout struct {
	layout
}

func CanvasLayout(bounds rl.Rectangle) canvasLayout {
	return canvasLayout{
		layout: layout{Bounds: bounds},
	}
}

func (cl *canvasLayout) Layout(anchor rl.Vector2, pivot rl.Vector2, rect rl.Rectangle) rl.Rectangle {
	anchorp := anchor.MultByVector(cl.Bounds.WH())
	pivotp := pivot.MultByVector(rect.WH())
	return rl.NewRectangleV(cl.Bounds.XY().Add(anchorp).Add(rect.XY()).Sub(pivotp), rect.WH())
}

type horizontalLayout struct {
	layout
	spacing  int
	position int
}

func HorizontalLayout(bounds rl.Rectangle, spacing int) horizontalLayout {
	return horizontalLayout{
		layout:   layout{Bounds: bounds},
		spacing:  spacing,
		position: 0,
	}
}

func (hl *horizontalLayout) Layout(wh rl.Vector2, justify Justyfy) rl.Rectangle {
	whY, dy := justify.Justyfy(wh.Y(), hl.Bounds.Height())
	r := rl.NewRectangle(hl.Bounds.X()+float32(hl.position), hl.Bounds.Y()+dy, wh.X(), whY)
	hl.position += int(wh.X()) + hl.spacing
	return r
}

func (hl *horizontalLayout) Fill(wh rl.Vector2, justify Justyfy) rl.Rectangle {
	whY, dy := justify.Justyfy(wh.Y(), hl.Bounds.Height())
	r := rl.NewRectangle(hl.Bounds.X()+float32(hl.position), hl.Bounds.Y()+dy, hl.Bounds.Width()-float32(hl.position), whY)
	hl.position = int(hl.Bounds.Width())
	return r
}

type verticalLayout struct {
	layout
	spacing  int
	position int
}

func VerticalLayout(bounds rl.Rectangle, spacing int) verticalLayout {
	return verticalLayout{
		layout:   layout{Bounds: bounds},
		spacing:  spacing,
		position: 0,
	}
}

func (vl *verticalLayout) Layout(wh rl.Vector2, justify Justyfy) rl.Rectangle {
	whX, dx := justify.Justyfy(wh.X(), vl.Bounds.Width())
	r := rl.NewRectangle(vl.Bounds.X()+dx, vl.Bounds.Y()+float32(vl.position), whX, wh.Y())
	vl.position += int(wh.Y()) + vl.spacing
	return r
}

func (vl *verticalLayout) Fill(wh rl.Vector2, justify Justyfy) rl.Rectangle {
	whX, dx := justify.Justyfy(wh.X(), vl.Bounds.Width())
	r := rl.NewRectangle(vl.Bounds.X()+dx, vl.Bounds.Y()+float32(vl.position), whX, vl.Bounds.Height()-float32(vl.position))
	vl.position = int(vl.Bounds.Height())
	return r
}
