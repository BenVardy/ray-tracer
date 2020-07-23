package mats

import "github.com/benvardy/raytracing/core"

type vector3 = core.Vector3

// Material stores information about a material
type Material struct {
	Ka, Kd, Ks   vector3
	Roughness    float64
	Reflectivity float64
}

var standardAmbient = vector3{0.1, 0.1, 0}

var WallMaterial = Material{
	Ka:           vector3{0.5, 0.5, 0.4},
	Kd:           vector3{0.3, 0.3, 0.3},
	Ks:           vector3{0.01, 0.01, 0.01},
	Roughness:    100,
	Reflectivity: 0.1,
}

var Ball1 = Material{
	Ka:           standardAmbient,
	Kd:           vector3{.8, 0, 0},
	Ks:           vector3{.05, .03, .03},
	Roughness:    0.8,
	Reflectivity: 0.1,
}

var Ball2 = Material{
	Ka:           standardAmbient,
	Kd:           vector3{0, .6, 0},
	Ks:           vector3{.03, .05, .03},
	Roughness:    0.8,
	Reflectivity: 0.1,
}

var Ball3 = Material{
	Ka:           standardAmbient,
	Kd:           vector3{0, 0, .8},
	Ks:           vector3{.01, .01, .03},
	Roughness:    0.8,
	Reflectivity: 0.1,
}
