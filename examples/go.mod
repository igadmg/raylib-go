module examples

go 1.21

replace github.com/gen2brain/raylib-go/raylib => ../raylib

replace github.com/gen2brain/raylib-go/raygui => ../raygui

replace github.com/gen2brain/raylib-go/easings => ../easings

replace github.com/gen2brain/raylib-go/physics => ../physics

replace github.com/EliCDavis/vector => ../vector

require (
	github.com/gen2brain/raylib-go/easings v0.0.0-20240227114648-c3665eb9abf8
	github.com/gen2brain/raylib-go/physics v0.0.0-20240227114648-c3665eb9abf8
	github.com/gen2brain/raylib-go/raygui v0.0.0-20240227114648-c3665eb9abf8
	github.com/gen2brain/raylib-go/raylib v0.0.0-20240227114648-c3665eb9abf8
	github.com/jakecoffman/cp v1.2.1
	github.com/neguse/go-box2d-lite v0.0.0-20170921151050-5d8ed9b7272b
)

require (
	github.com/EliCDavis/vector v1.6.0 // indirect
	github.com/ebitengine/purego v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)
