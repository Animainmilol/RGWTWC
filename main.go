package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DefaultScreenWidth  = 800
	DefaultScreenHeight = 450
	TargetFPS           = 120

	MaxZoom = 10.0
	MinZoom = 0.8

	DefaultZoom          = 1.0
	DefaultRotation      = 0.0
	DefaultRotationSpeed = 5.0
	DefaultZoomSpeed     = 5.0
	DefaultFollowSpeed   = 5.0
)

type CameraController struct {
	camera rl.Camera2D

	ActiveEffects []CameraEffect

	RotationSpeed float32
	ZoomSpeed     float32
	FollowSpeed   float32

	followTarget rl.Vector2
}

type CameraEffect struct {
	Name string
	ZoomFactor float32
	Rotation float32
}

func NewCameraController() *CameraController {
	offsetX := float32(DefaultScreenWidth) / 2
	offsetY := float32(DefaultScreenHeight) / 2

	baseEffect := CameraEffect{
        Name: "UserControl",
        ZoomFactor: DefaultZoom,
        Rotation: DefaultRotation,
    }

	return &CameraController{
		camera: rl.NewCamera2D(
			rl.Vector2{X: offsetX, Y: offsetY},
			rl.Vector2{X: 0, Y: 0},
			DefaultRotation,
			DefaultZoom,
		),

		ActiveEffects: []CameraEffect{baseEffect},

		RotationSpeed: DefaultRotationSpeed,
		ZoomSpeed:     DefaultZoomSpeed,
		FollowSpeed:   DefaultFollowSpeed,

		followTarget: rl.Vector2{X: 0, Y: 0},
	}
}

func NewCameraEffect(name string, zoomFactor, rotation float32) *CameraEffect {
	return &CameraEffect{
		Name: name,
		ZoomFactor: zoomFactor,
		Rotation: rotation,
	}
}

func (cc *CameraController) GetCamera() rl.Camera2D {
	return cc.camera
}

func (cc *CameraController) SetFollowTarget(target rl.Vector2) {
    cc.followTarget = target
}

func (cc *CameraController) Update() {
	var targetZoom float32
	var targetRotation float32

	for _, v := range cc.ActiveEffects {
		targetZoom *= v.ZoomFactor
	}
	for _, v := range cc.ActiveEffects {
		targetRotation += v.Rotation
	}

	cc.camera.Zoom = rl.Lerp(cc.camera.Zoom, targetZoom, cc.ZoomSpeed*rl.GetFrameTime())
	cc.camera.Rotation = rl.Lerp(cc.camera.Rotation, targetRotation, cc.RotationSpeed*rl.GetFrameTime())
	cc.camera.Target = rl.Vector2Lerp(cc.camera.Target, cc.followTarget, cc.FollowSpeed*rl.GetFrameTime())
}

func main() {
	rl.InitWindow(DefaultScreenWidth, DefaultScreenHeight, "RGWTWC")
	defer rl.CloseWindow()

	rl.SetTargetFPS(TargetFPS)

	cc := NewCameraController()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		cc.Update()
		rl.BeginMode2D(cc.GetCamera())

		rl.ClearBackground(rl.Black)
		rl.DrawText("yoo", 0, 0, 30, rl.RayWhite)

		rl.EndDrawing()
	}
}
