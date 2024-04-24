module github.com/gen2brain/raylib-go/raylib

go 1.21

replace github.com/EliCDavis/vector => ../vector

require (
	github.com/EliCDavis/vector v1.6.0
	github.com/ebitengine/purego v0.7.0
	golang.org/x/sys v0.19.0
)

require golang.org/x/exp v0.0.0-20240409090435-93d18d7e34b8 // indirect
