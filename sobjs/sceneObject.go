package sobjs

import (
	"github.com/benvardy/raytracing/core"
	"github.com/benvardy/raytracing/mats"
)

// vector3 is an aillis for core.Vector3 for ease of use
type vector3 = core.Vector3

// SceneObject is the interface that describes any object that can be traced in the ray tracer
type SceneObject interface {
	// IntersectWithRay takes a line in the form of: s - a start vector, and d - a direction vector
	IntersectWithRay(s, d core.Vector3) *core.Vector3
	// Gets the normal of the object at position p and adjusts based on the position of the light l
	// This adjustment allows objects like Plane and disk to be visible from below
	GetNormal(p core.Vector3, l core.Vector3) core.Vector3
	GetMaterial() mats.Material
}

// type Cylinder struct {
// 	Centre core.Vector3
// 	Mat    mats.Material

// 	Normal core.Vector3
// 	// Total height of the cylinder
// 	Height float64
// 	Radius float64

// 	// Store the top and bottom planes to interest with
// 	topDisk    *Disk
// 	bottomDisk *Disk
// }

// func NewCylinder(centre, normal core.Vector3, height, radius float64, mats.Material mats.Material) *Cylinder {
// 	n := normal.Normalize()
// 	topDisk := NewDisk(centre.Add(n.Smult(height/2)), n, radius, mats.Material)
// 	bottomDisk := NewDisk(centre.Add(n.Smult(-height/2)), n, radius, mats.Material)

// 	return &Cylinder{centre, mats.Material, normal, height / 2, radius, topDisk, bottomDisk}
// }

// func (cylinder *Cylinder) IntersectWithRay(s, d core.Vector3) *core.Vector3 {
// 	d = d.Normalize()
// 	c, n, r := cylinder.Centre.AsSlice(), cylinder.Normal, cylinder.Radius
// 	p1 := s.AsSlice()

// 	d3 := d.Cross(n).Normalize()

// 	// Define the transpose to get the vectors as slices in one place
// 	mT := [][]float64{d.AsSlice(), n.AsSlice(), d3.AsSlice()}

// 	// Three for number of variables
// 	dim := 3
// 	m := make([][]float64, dim)
// 	v := make([]float64, dim)
// 	for i := 0; i < dim; i++ {
// 		v[i] = c[i] - p1[i]
// 		m[i] = make([]float64, dim)

// 		for j := 0; j < dim; j++ {
// 			m[i][j] = mT[j][i]
// 		}
// 	}

// 	lambda := SolveGauss(m, v)

// 	if math.Abs(lambda[2]) > r {
// 		return nil
// 	}

// 	if math.Abs(lambda[1]) > cylinder.Height {
// 		if inSect := cylinder.bottomDisk.IntersectWithRay(s, d); inSect != nil {
// 			return inSect
// 		}

// 		return cylinder.topDisk.IntersectWithRay(s, d)
// 	}

// 	newPointD := s.Add(n.Smult(lambda[0] - math.Sqrt(r*r-lambda[2]*lambda[2])))
// 	return &newPointD
// }
