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
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(DefaultScreenWidth, DefaultScreenHeight, "RGWTWC")
	defer rl.CloseWindow()

	rl.SetTargetFPS(TargetFPS)

	cc := NewCameraController()

	for !rl.WindowShouldClose() {
		cc.GetCamera().Offset.X = float32(rl.GetRenderWidth() / 2)
		cc.GetCamera().Offset.Y = float32(rl.GetRenderHeight() / 2)
		cc.Update()

		rl.BeginDrawing()

		rl.BeginMode2D(*cc.GetCamera())

		rl.ClearBackground(rl.Black)
		rl.DrawText("yoo", 0, 0, 30, rl.RayWhite)

		rl.EndDrawing()
	}
}
