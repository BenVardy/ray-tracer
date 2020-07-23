package tracer

import (
	"github.com/benvardy/raytracing/core"
	"github.com/benvardy/raytracing/sobjs"
)

// Scene represents the 3D space that we are ray tracing
type Scene struct {
	upDirection   core.Vector3
	leftDirection core.Vector3
	lookDirection core.Vector3
	eyePosition   core.Vector3
	meshTopLeft   core.Vector3

	meshDistance float64
	pixelWidth   float64

	focalDistance float64
	apertureSize  float64

	ScreenWidth  int
	ScreenHeight int

	Objects []sobjs.SceneObject
	Lights  []*core.SceneLight

	Ia core.Vector3
}

// NewScene creates a scene
func NewScene(leftDirection, lookDirection, eyePosition core.Vector3, meshDistance, pixelWidth, focalDistance, apertureSize float64, screenWidth, screenHeight int, ia core.Vector3) *Scene {
	upDirection := lookDirection.Cross(leftDirection).Normalize()

	meshTopLeft := eyePosition.Add(lookDirection.Smult(meshDistance)).Add(leftDirection.Smult(pixelWidth * float64(screenWidth) / 2.0)).Add(upDirection.Smult(pixelWidth * float64(screenHeight) / 2.0))

	return &Scene{
		upDirection,
		leftDirection,
		lookDirection,
		eyePosition,
		meshTopLeft,
		meshDistance,
		pixelWidth,
		focalDistance,
		apertureSize,
		screenWidth,
		screenHeight,
		make([]sobjs.SceneObject, 0),
		make([]*core.SceneLight, 0),
		ia,
	}
}

// GetEye returns the eyePosition
func (s *Scene) GetEye() core.Vector3 {
	return s.eyePosition
}

// GetRayToMesh returns a normal vector from the eye to the (x, y) position on the mesh
func (s *Scene) GetRayToMesh(x, y int) core.Vector3 {
	down := s.upDirection.Smult(-1)
	right := s.leftDirection.Smult(-1)

	pixelPosition := s.meshTopLeft.Add(right.Smult(float64(x) * s.pixelWidth)).Add(down.Smult(float64(y) * s.pixelWidth))

	return pixelPosition.Subtract(s.eyePosition).Normalize()
}

// AddSceneObject adds an object to the scene to be rendered
func (s *Scene) AddSceneObject(obj sobjs.SceneObject) {
	s.Objects = append(s.Objects, obj)
}

// AddSceneLight adds a light to the scene
func (s *Scene) AddSceneLight(light *core.SceneLight) {
	s.Lights = append(s.Lights, light)
}
