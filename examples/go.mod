module examples

go 1.24

toolchain go1.24.2

replace github.com/igadmg/raylib-go/raylib => ../raylib

replace github.com/igadmg/raylib-go/raygui => ../raygui

replace github.com/igadmg/raylib-go/easings => ../easings

replace github.com/igadmg/raylib-go/physics => ../physics

replace github.com/igadmg/gamemath => ../../gamemath

require (
	github.com/igadmg/gamemath v1.6.0
	github.com/igadmg/goex v0.0.0-20250407220752-712c023573b8
	github.com/igadmg/raylib-go/easings v0.0.0-20250327112125-46827dd07487
	github.com/igadmg/raylib-go/physics v0.0.0-20250327112125-46827dd07487
	github.com/igadmg/raylib-go/raygui v0.0.0-20250327112125-46827dd07487
	github.com/igadmg/raylib-go/raylib v0.0.0-20250327112125-46827dd07487
	github.com/jakecoffman/cp v1.2.1
	github.com/neguse/go-box2d-lite v0.0.0-20170921151050-5d8ed9b7272b
)

require (
	github.com/chewxy/math32 v1.11.1 // indirect
	github.com/ebitengine/purego v0.8.2 // indirect
	golang.org/x/exp v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
