package sobjs

import (
	"math"

	"github.com/benvardy/raytracing/mats"
)

// Sphere is a sphere SceneObject
type Sphere struct {
	Position vector3
	Mat      mats.Material
	Radius   float64
}

// NewSphere constructs a new sphere sceneObject
func NewSphere(pos vector3, r float64, material mats.Material) *Sphere {
	return &Sphere{pos, material, r}
}

// IntersectWithRay implements the SceneObject function
//  Solves the equation `aλ^2 + bλ + c = 0`, where:
//  `a = d.d, b = 2*(d.(s - c)), c = (s - c).(s - c) - r^2`
func (sphere *Sphere) IntersectWithRay(s, d vector3) *vector3 {
	center := sphere.Position
	r := sphere.Radius

	a := d.Dot(d)
	// Check if d is a 0 vector and if so the ray is invalid
	if a == 0 {
		return nil
	}

	// offset is the realigned center to minimize the number of times s - center is done
	// for calculating b and c
	offset := s.Subtract(center)

	b := 2 * d.Dot(offset)
	c := offset.Dot(offset) - r*r

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	// lambda1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	// lambda2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	// p1 := s.Add(d.Smult(lambda1))
	// p2 := s.Add(d.Smult(lambda2))

	// if p1.Length() < p2.Length() {
	// 	return &p1
	// }

	// return &p2

	// Only have to calculate the smaller value of lambda as
	lambda := (-b - math.Sqrt(discriminant)) / (2 * a)
	if lambda < 0 {
		return nil
	}

	p := s.Add(d.Smult(lambda))

	return &p
}

// GetNormal gets the normal at the point p
func (sphere *Sphere) GetNormal(p, _ vector3) vector3 {
	return p.Subtract(sphere.Position).Normalize()
}

// GetMaterial gets the mats.Material
func (sphere *Sphere) GetMaterial() mats.Material {
	return sphere.Mat
}
