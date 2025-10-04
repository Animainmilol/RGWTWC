package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DefaultScreenWidth  = 800
	DefaultScreenHeight = 450
	TargetFPS           = 120
)

func main() {
	rl.InitWindow(DefaultScreenWidth, DefaultScreenHeight, "RGWTWC")
	rl.SetConfigFlags(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(TargetFPS)

	cc := NewCameraController()

	for !rl.WindowShouldClose() {
		cc.Update()

		rl.BeginDrawing()

		rl.BeginMode2D(cc.GetCamera())

		rl.ClearBackground(rl.Black)
		rl.DrawText("yoo", 0, 0, 30, rl.RayWhite)

		rl.EndDrawing()
	}
}
