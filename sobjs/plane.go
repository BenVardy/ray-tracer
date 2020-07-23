package sobjs

import "github.com/benvardy/raytracing/mats"

// Plane is a plane SceneObject
type Plane struct {
	Position vector3
	Mat      mats.Material
	Normal   vector3
}

// NewPlane creates a new Plane should be used to normalize the normal vector
func NewPlane(pos, normal vector3, material mats.Material) *Plane {
	return &Plane{pos, material, normal.Normalize()}
}

// IntersectWithRay implements the SceneObject function
//  Solves the equation `(s + λd - q) . n = 0`
func (plane *Plane) IntersectWithRay(s, d vector3) *vector3 {
	n := plane.Normal
	pos := plane.Position

	if d.Dot(n) == 0 {
		return nil
	}

	lambda := n.Dot(pos.Subtract(s)) / n.Dot(d)

	if lambda <= 0 {
		return nil
	}

	res := s.Add(d.Smult(lambda))
	return &res
}

// GetNormal gets the normal at the point p
func (plane *Plane) GetNormal(p, _ vector3) vector3 {
	n := plane.Normal.Normalize()

	return n
}

// GetMaterial gets the mats.Material
func (plane *Plane) GetMaterial() mats.Material {
	return plane.Mat
}
