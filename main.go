package main

import (
	"flag"
	"fmt"
	"math"
	"time"

	"github.com/benvardy/raytracing/core"
	"github.com/benvardy/raytracing/mats"
	"github.com/benvardy/raytracing/sobjs"
	"github.com/benvardy/raytracing/tracer"
)

type vector3 = core.Vector3

func printTimeTaken(lab string, start time.Time) {
	fmt.Printf("TIMER: %s : %v\n", lab, time.Since(start))
}

func main() {
	// Flags
	var saveLoc string
	flag.StringVar(&saveLoc, "save", "image.png", "The path of the image save")
	flag.StringVar(&saveLoc, "s", "image.png", "The path of the image save (shorthand)")

	var width, height int
	flag.IntVar(&width, "w", 1920, "The width of the image")
	flag.IntVar(&height, "h", 1080, "The width of the image")

	var dof, nshadows bool
	flag.BoolVar(&dof, "dof", false, "Toggle Depth of Field")
	flag.BoolVar(&nshadows, "ns", false, "Toggle nice shadows")
	flag.Parse()

	defer printTimeTaken("Ray Trace", time.Now())

	img := core.NewImage(width, height)

	base := 1920.0 * 1080.0
	pixelWidth := 0.2 * math.Log2(base/float64(width*height))
	// To fix if with and height are 1920*1080
	if pixelWidth == 0 {
		pixelWidth = 0.2
	}

	scene := tracer.NewScene(
		vector3{-1, 0, 0}, // left
		vector3{0, 1, 0},  // look
		vector3{0, 0, 0},  // eye
		150,               // grid dist
		pixelWidth,        // pixel width
		50,                // focal dist
		0.6,               // aperture size
		img.Width,
		img.Height,
		vector3{0.05, 0.05, 0.05}, // ambient
	)

	// Objects
	// Sphere
	scene.AddSceneObject(sobjs.NewSphere(vector3{10, 50, 5}, 10, mats.Ball1))
	scene.AddSceneObject(sobjs.NewSphere(vector3{-2.5, 25, 0}, 5, mats.Ball2))
	scene.AddSceneObject(sobjs.NewSphere(vector3{10, 27, -2.5}, 2.5, mats.Ball3))
	scene.AddSceneObject(sobjs.NewSphere(vector3{0, -1, -2.5}, 2.5, mats.Ball3))

	// Floor
	scene.AddSceneObject(sobjs.NewPlane(vector3{0, 0, -5}, vector3{0, 0, 1}, mats.WallMaterial))
	// Walls
	scene.AddSceneObject(sobjs.NewPlane(vector3{0, 1000, 0}, vector3{0, -1, 0}, mats.WallMaterial))

	// Lights
	scene.AddSceneLight(core.NewSceneLight(vector3{4.5, 26, -4}, vector3{.6, .6, .6}, 1))
	// Studio Lights
	scene.AddSceneLight(core.NewSceneLight(vector3{100, -100, 30}, vector3{.3, .3, .3}, 1))
	scene.AddSceneLight(core.NewSceneLight(vector3{-100, -100, 30}, vector3{0.3, .3, 0.3}, 1))
	scene.AddSceneLight(core.NewSceneLight(vector3{100, 100, 30}, vector3{0.3, 0.3, .3}, 1))
	scene.AddSceneLight(core.NewSceneLight(vector3{-100, 100, 30}, vector3{.3, .3, .3}, 1))

	tracer.Trace(scene, img, dof, nshadows)

	img.PrintToFile(saveLoc)
}
