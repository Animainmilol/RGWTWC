package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxZoom float32 = 10.0
	MinZoom float32 = 0.8

	DefaultZoom          float32 = 1.0
	DefaultRotation      float32 = 0.0
	DefaultRotationSpeed float32 = 5.0
	DefaultZoomSpeed     float32 = 5.0
	DefaultFollowSpeed   float32 = 5.0
)

type MovementMode int

const (
	ModeInstant MovementMode = iota
	ModeLerp
)

type CameraController struct {
	camera rl.Camera2D

	activeEffects []*CameraEffect

	RotationSpeed float32
	ZoomSpeed     float32
	FollowSpeed   float32

	followTarget rl.Vector2

	ZoomMode     MovementMode
	RotationMode MovementMode
	FollowMode   MovementMode
}

type CameraEffect struct {
	Name       string
	ZoomFactor float32
	Rotation   float32
}

func NewCameraController() *CameraController {
	offsetX := rl.GetRenderWidth() / 2.0
	offsetY := rl.GetRenderHeight() / 2.0

	baseEffect := &CameraEffect{
		Name:       "UserControl",
		ZoomFactor: DefaultZoom,
		Rotation:   DefaultRotation,
	}

	return &CameraController{
		camera: rl.NewCamera2D(
			rl.Vector2{X: float32(offsetX), Y: float32(offsetY)},
			rl.Vector2{X: 0, Y: 0},
			DefaultRotation,
			DefaultZoom,
		),

		activeEffects: []*CameraEffect{baseEffect},

		RotationSpeed: DefaultRotationSpeed,
		ZoomSpeed:     DefaultZoomSpeed,
		FollowSpeed:   DefaultFollowSpeed,

		followTarget: rl.Vector2{X: 0, Y: 0},

		ZoomMode:     ModeLerp,
		RotationMode: ModeLerp,
		FollowMode:   ModeLerp,
	}
}

func NewCameraEffect(name string, zoomFactor, rotation float32) *CameraEffect {
	return &CameraEffect{
		Name:       name,
		ZoomFactor: zoomFactor,
		Rotation:   rotation,
	}
}

func (cc *CameraController) AddEffect(ce *CameraEffect) {
	for _, effect := range cc.activeEffects {
		if effect.Name == ce.Name {
			return
		}
	}
	cc.activeEffects = append(cc.activeEffects, ce)
}

func (cc *CameraController) RemoveEffect(name string) {
	for i, effect := range cc.activeEffects {
		if effect.Name == name {
			lastIndex := len(cc.activeEffects) - 1
			cc.activeEffects[i] = cc.activeEffects[lastIndex]
			cc.activeEffects = cc.activeEffects[:lastIndex]
			return
		}
	}
}

func (cc *CameraController) GetCamera() rl.Camera2D {
	return cc.camera
}

func (cc *CameraController) SetFollowTarget(target rl.Vector2) {
	cc.followTarget = target
}

func (cc *CameraController) Update() {
	targetZoom := DefaultZoom
	targetRotation := DefaultRotation

	for _, v := range cc.activeEffects {
		targetZoom *= v.ZoomFactor
		targetRotation += v.Rotation
	}

	targetZoom = rl.Clamp(targetZoom, MinZoom, MaxZoom)
	
	switch cc.ZoomMode{
	case ModeInstant:
		cc.camera.Zoom = targetZoom
	case ModeLerp:
		cc.camera.Zoom = rl.Lerp(cc.camera.Zoom, targetZoom, cc.ZoomSpeed*rl.GetFrameTime())
	}

	switch cc.RotationMode{
	case ModeInstant:
		cc.camera.Rotation = targetRotation
	case ModeLerp:
		cc.camera.Rotation = rl.Lerp(cc.camera.Rotation, targetRotation, cc.RotationSpeed*rl.GetFrameTime())
	}

	switch cc.FollowMode{
	case ModeInstant:
		cc.camera.Target = cc.followTarget
	case ModeLerp:
		cc.camera.Target = rl.Vector2Lerp(cc.camera.Target, cc.followTarget, cc.FollowSpeed*rl.GetFrameTime())
	}
}