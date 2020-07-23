package tracer

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"strings"

	"github.com/benvardy/raytracing/core"
	"github.com/benvardy/raytracing/sobjs"
)

type vector3 = core.Vector3

// Trace implements a basic ray tracer
func Trace(scene *Scene, img *core.Image, DOF, shading bool) {

	maxPos := 25
	apertureSize := scene.apertureSize
	if !DOF {
		maxPos = 1
		apertureSize = 0
	}

	totalProg := float64(scene.ScreenHeight * scene.ScreenWidth)
	totalHashes := 50

	for y := 0; y < scene.ScreenHeight; y++ {
		for x := 0; x < scene.ScreenWidth; x++ {
			// focal point
			d := scene.GetRayToMesh(x, y).Normalize()
			P := scene.GetEye().Add(d.Smult(scene.focalDistance))

			// Changed to one to temp stop
			var c vector3
			if maxPos > 1 {
				for i := 0; i < maxPos; i++ {
					leftMod := scene.leftDirection.Smult(rand.Float64() - 0.5).Smult(apertureSize)
					upMod := scene.upDirection.Smult(rand.Float64() - 0.5).Smult(apertureSize)

					newEye := scene.GetEye().Add(leftMod).Add(upMod)
					c = c.Add(findColor(scene, newEye, P.Subtract(newEye).Normalize(), 0, nil, shading))
				}
				c = c.Smult(1.0 / float64(maxPos))
			} else {
				c = findColor(scene, scene.eyePosition, d, 0, nil, shading)
			}

			// Gamma
			gamma := 2.2
			c.X = math.Pow(c.X, 1/gamma)
			c.Y = math.Pow(c.Y, 1/gamma)
			c.Z = math.Pow(c.Z, 1/gamma)

			// Convert to RGB
			c.X = math.Min(255, c.X*255)
			c.Y = math.Min(255, c.Y*255)
			c.Z = math.Min(255, c.Z*255)

			img.SetPixel(x, y, color.RGBA{uint8(c.X), uint8(c.Y), uint8(c.Z), 0xff})

			doneNow := float64(y*scene.ScreenWidth + x)
			noHash := int(math.Ceil(float64(totalHashes) * doneNow / totalProg))

			fmt.Printf("[%s%s]\r", strings.Repeat("#", noHash), strings.Repeat(".", totalHashes-noHash))
		}
	}
	fmt.Println()
}

func findColor(scene *Scene, s, d vector3, depth float64, prevObj sobjs.SceneObject, shading bool) vector3 {
	// Distributed shading
	maxTotalHit := 25
	if !shading {
		maxTotalHit = 1
	}

	background := vector3{}

	if depth >= 3 {
		return background
	}

	var closestPos *vector3
	var closestObject sobjs.SceneObject

	for _, o := range scene.Objects {
		posOfIntPtr := o.IntersectWithRay(s, d)

		if posOfIntPtr != nil {
			if closestObject == nil || posOfIntPtr.Subtract(s).Length() < closestPos.Subtract(s).Length() {
				closestPos = posOfIntPtr
				closestObject = o
			}
		}
	}

	if prevObj != nil && closestObject == prevObj {
		return background
	}

	if closestObject != nil {
		material := closestObject.GetMaterial()

		var reflectedIntensity vector3
		I := vector3{}

		if material.Reflectivity > 0 {
			inN := closestObject.GetNormal(*closestPos, d).Normalize()
			mirrorDir := inN.Smult(d.Dot(inN)).Add(d).Smult(-2)

			reflectedIntensity = findColor(scene, *closestPos, mirrorDir, depth+1, closestObject, shading)
		}

		// We saw an object

		I.X += scene.Ia.X * material.Ka.X
		I.Y += scene.Ia.Y * material.Ka.Y
		I.Z += scene.Ia.Z * material.Ka.Z

		for _, light := range scene.Lights {
			size := light.Size
			if !shading {
				size = 0
			}

			totalHit := 0

			IL := vector3{}
			N := closestObject.GetNormal(*closestPos, light.Position).Normalize()

			if maxTotalHit > 1 {
				for i := 0; i < maxTotalHit; i++ {

					leftMod := scene.leftDirection.Smult(rand.Float64() - 0.5).Smult(size)
					lookMod := scene.lookDirection.Smult(rand.Float64() - 0.5).Smult(size)

					LPos := light.Position.Add(leftMod).Add(lookMod)

					L := LPos.Subtract(*closestPos).Normalize()

					visible := true
					for _, o := range scene.Objects {
						if o != closestObject {
							shadingObjPos := o.IntersectWithRay(*closestPos, L)

							if shadingObjPos != nil && (shadingObjPos.Subtract(*closestPos)).Dot(L) > 0 && shadingObjPos.Subtract(*closestPos).Length() < LPos.Subtract(*closestPos).Length() {
								visible = false
								break
							}
						}
					}

					if visible {
						totalHit++
					}

				}

				L := light.Position.Subtract(*closestPos).Normalize()

				// Diffuse I_d = I_l * k_d * (N.L)
				if dot := N.Dot(L); dot > 0 {
					IL.X += light.Intensity.X * material.Kd.X * dot
					IL.Z += light.Intensity.Z * material.Kd.Z * dot
					IL.Y += light.Intensity.Y * material.Kd.Y * dot

					// Specular
					V := scene.GetEye().Subtract(*closestPos).Normalize()
					R := N.Smult(2 * L.Dot(N)).Subtract(L).Normalize()
					if R.Dot(V) > 0 {
						dotN := math.Pow(R.Dot(V), material.Roughness)
						IL.X += light.Intensity.X * material.Ks.X * dotN
						IL.Y += light.Intensity.Y * material.Ks.Y * dotN
						IL.Z += light.Intensity.Z * material.Ks.Z * dotN
					}
				}

				I = I.Add(IL.Smult(float64(totalHit) / float64(maxTotalHit)))

			} else {
				L := light.Position.Subtract(*closestPos).Normalize()

				visible := true
				for _, o := range scene.Objects {
					if o != closestObject {
						shadingObjPos := o.IntersectWithRay(*closestPos, L)

						if shadingObjPos != nil && (shadingObjPos.Subtract(*closestPos)).Dot(L) > 0 && shadingObjPos.Subtract(*closestPos).Length() < light.Position.Subtract(*closestPos).Length() {
							visible = false
							break
						}
					}
				}

				// Diffuse I_d = I_l * k_d * (N.L)
				if dot := N.Dot(L); visible && dot > 0 {
					IL.X += light.Intensity.X * material.Kd.X * dot
					IL.Z += light.Intensity.Z * material.Kd.Z * dot
					IL.Y += light.Intensity.Y * material.Kd.Y * dot

					// Specular
					V := scene.GetEye().Subtract(*closestPos).Normalize()
					R := N.Smult(2 * L.Dot(N)).Subtract(L).Normalize()
					if R.Dot(V) > 0 {
						dotN := math.Pow(R.Dot(V), material.Roughness)
						IL.X += light.Intensity.X * material.Ks.X * dotN
						IL.Y += light.Intensity.Y * material.Ks.Y * dotN
						IL.Z += light.Intensity.Z * material.Ks.Z * dotN
					}
				}
				I = I.Add(IL)

			}

		}

		return reflectedIntensity.Smult(material.Reflectivity).Add(I.Smult(1 - material.Reflectivity))
	}

	// Black
	return background
}
