package sobjs

import "github.com/benvardy/raytracing/mats"

// Disk is a disk - a plane with bounds SceneObject
type Disk struct {
	RootPlane *Plane
	Radius    float64
}

// NewDisk creates a new disk from pos, a normal, and a radius
func NewDisk(pos vector3, normal vector3, radius float64, material mats.Material) *Disk {
	return &Disk{NewPlane(pos, normal, material), radius}
}

// IntersectWithRay implements the SceneObject function
func (disk *Disk) IntersectWithRay(s, d vector3) *vector3 {
	rp := disk.RootPlane

	vPtr := rp.IntersectWithRay(s, d)
	if vPtr != nil && rp.Position.Subtract(*vPtr).Length() <= disk.Radius {
		return vPtr
	}

	return nil
}

// GetNormal gets the normal at the point p
func (disk *Disk) GetNormal(p, l vector3) vector3 {
	n := disk.RootPlane.GetNormal(p, l)
	if n.Dot(l.Subtract(p)) < 0 {
		return n.Smult(-1)
	}
	return n
}

// GetMaterial gets the mats.Material
func (disk *Disk) GetMaterial() mats.Material {
	return disk.RootPlane.GetMaterial()
}
