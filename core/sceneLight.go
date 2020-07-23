package core

type SceneLight struct {
	Position  Vector3
	Intensity Vector3
	Size      float64
}

func NewSceneLight(pos, intensity Vector3, size float64) *SceneLight {
	return &SceneLight{pos, intensity, size}
}
