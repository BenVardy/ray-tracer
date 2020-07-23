package sobjs

import (
	"math"
	"math/cmplx"

	"github.com/benvardy/raytracing/mats"
)

func solveCubic(a, b, c float64) (complex128, complex128, complex128) {
	var alpha, beta, gamma complex128
	if c == 0 {
		alpha = 0.0 + 0.0i
		ca, cb := complex(a, 0.0), complex(b, 0.0)
		beta = (-ca + cmplx.Sqrt(ca*ca-4.0*cb)) / 2.0
		gamma = (-ca - cmplx.Sqrt(ca*ca-4.0*cb)) / 2.0
	} else {
		p := complex((3.0*b-a*a)/3.0, 0.0)
		q := complex((9*b*a-27*c-2*math.Pow(a, 3))/27.0, 0.0)

		tCubed := (q + cmplx.Sqrt(q*q+4.0*(cmplx.Pow(p, 3.0)/27.0))) / 2.0
		t := cmplx.Pow(tCubed, 1.0/3.0)

		alpha = t - p/(3.0*t) - complex(a/3.0, 0.0)

		newB := complex(a, 0.0)*alpha + alpha*alpha
		beta = (-newB + cmplx.Sqrt(newB*newB-4.0*alpha*complex(-c, 0.0))) / (2.0 * alpha)
		gamma = (-newB - cmplx.Sqrt(newB*newB-4.0*alpha*complex(-c, 0.0))) / (2.0 * alpha)
	}

	return alpha, beta, gamma
}

// SolveQuartic https://en.wikipedia.org/wiki/Quartic_function#Ferrari's_solution
func SolveQuartic(a, b, c, d float64) (complex128, complex128, complex128, complex128) {
	var alpha, beta, gamma, delta complex128
	if d == 0 {
		alpha = 0
		beta, gamma, delta = solveCubic(a, b, c)
	} else {
		p := (8*b - 3*a*a) / 8
		q := (math.Pow(a, 3) - 4*a*b + 8*c) / 8
		r := (-3*math.Pow(a, 4) + 256*d - 64*a*c + 16*a*a*b) / 256

		// solve 8m^3 + 8pm^2 + (2p^2 - 8r)m - q^2 = 0

		thing := (2*p*p - 8*r)
		_ = thing
		m1, m2, m3 := solveCubic(p, (2*p*p-8*r)/8, -(q*q)/8)
		roots := []complex128{m1, m2, m3}
		var m complex128
		// Loop through the values of m and pick one that is non-0
		for _, mi := range roots {
			if mi != 0 {
				m = mi
				break
			}
		}

		cp, cq, ca := complex(p, 0), complex(q, 0), complex(a, 0)
		sqrtM := cmplx.Sqrt(m)

		alpha = -ca/4 + (math.Sqrt2*sqrtM+cmplx.Sqrt(-(2*cp+2*m+(math.Sqrt2*cq)/sqrtM)))/2
		beta = -ca/4 + (-math.Sqrt2*sqrtM+cmplx.Sqrt(-(2*cp+2*m-(math.Sqrt2*cq)/sqrtM)))/2
		gamma = -ca/4 + (math.Sqrt2*sqrtM-cmplx.Sqrt(-(2*cp+2*m+(math.Sqrt2*cq)/sqrtM)))/2
		delta = -ca/4 + (-math.Sqrt2*sqrtM-cmplx.Sqrt(-(2*cp+2*m-(math.Sqrt2*cq)/sqrtM)))/2
	}

	return alpha, beta, gamma, delta
}

type Torus struct {
	Position vector3
	Mat      mats.Material
	BigR     float64
	LittleR  float64
}

func NewTorus(pos vector3, R, r float64, mat mats.Material) *Torus {
	return &Torus{pos, mat, R, r}
}

func (torus *Torus) IntersectWithRay(s, d vector3) *vector3 {
	newS := torus.Position.Subtract(s)

	R, r := torus.BigR, torus.LittleR
	a := math.Pow(d.X, 4) + math.Pow(d.Y, 4) + math.Pow(d.Z, 4) + 2*d.X*d.X*d.Y*d.Y + 2*d.X*d.X*d.Z*d.Z + 2*d.Y*d.Y*d.Z*d.Z
	b := 4*math.Pow(d.X, 3)*s.X + 4*math.Pow(d.Y, 3)*s.Y + 4*math.Pow(d.Z, 3)*s.Z + 4*d.X*d.X*d.Y*s.Y + 4*d.X*d.X*d.Z*s.Z + 4*d.X*s.X*d.Y*d.Y + 4*d.Y*d.Y*d.Z*s.Z + 4*d.X*s.X*s.Z*s.Z + 4*d.Y*s.Y*s.Z*s.Z
	c := -2*R*R*d.X*d.X - 2*R*R*d.Y*d.Y - 2*R*R*d.Z*d.Z - 2*r*r*d.X*d.X - 2*r*r*d.Y*d.Y - 2*r*r*d.Z*d.Z + 6*d.X*d.X*s.X*s.X + 2*s.X*s.X*d.Y*d.Y + 8*d.X*s.X*d.Y*s.Y + 2*d.X*d.X*s.Y*s.Y + 6*d.Y*d.Y*s.Y*s.Y + 2*s.X*s.X*d.Z*d.Z + 2*s.Y*s.Y*d.Z*d.Z + 8*d.X*s.X*d.Z*s.Z + 8*d.Y*s.Y*d.Z*s.Z + 2*d.X*d.X*d.Z*s.Z + 2*d.Y*d.Y*s.Z*s.Z + 6*d.Z*d.Z*s.Z*s.Z
	e := -4*R*R*d.X*s.X - 4*R*R*d.Y*s.X + 4*R*R*d.X*s.X - 4*r*r*d.X*s.X - 4*r*r*d.Y*s.Y - 4*r*r*d.Z*s.Z + 4*d.X*s.X*s.X*s.X + 4*s.X*s.X*d.Y*s.Y + 4*d.X*s.X*s.Y*s.Y + 4*d.Y*s.Y*s.Y*s.Y + 4*s.X*s.X*d.Z*s.Z + 4*s.Y*s.Y*d.Z*s.Z + 4*s.Y*s.Y*d.Z*s.Z + 4*d.X*s.X*s.Z*s.Z + 4*d.Z*s.Z*s.Z*s.Z
	f := math.Pow(R, 4) - 2*R*R*s.X*s.X - 2*R*R*s.Y*s.Y + 2*R*R*s.Z*s.Z + math.Pow(r, 4) - 2*r*r*R*R - 2*r*r*s.X*s.X - 2*r*r*s.Y*s.Y - 2*r*r*s.Z*s.Z + math.Pow(s.X, 4) + math.Pow(s.Y, 4) + math.Pow(s.Z, 4) + 2*s.X*s.X*s.Y*s.Y + 2*s.X*s.X*s.Z*s.Z + 2*s.Y*s.Y*s.Z*s.Z

	alf, bet, gam, del := SolveQuartic(b/a, c/a, e/a, f/a)

	roots := []complex128{alf, bet, gam, del}
	reRoots := make([]float64, 0)

	for _, root := range roots {
		// *100000 for dec places
		if math.Round(imag(root)*100000) == 0 && real(root) > 0 {
			reRoots = append(reRoots, real(root))
		}
	}

	if len(reRoots) == 0 {
		return nil
	}

	smallest := math.MaxFloat64
	for _, root := range reRoots {
		if root < smallest {
			smallest = root
		}
	}

	pos := newS.Add(s.Smult(smallest))
	return &pos
}

func (torus *Torus) GetNormal(p, _ vector3) vector3 {
	a, b, c := torus.Position.X, torus.Position.Y, torus.Position.Z
	R := torus.BigR
	x := -4 * (math.Pow(a, 3) - 2*a*a*p.X + a*(b*b-2*b*p.Y+c*c-2*c*p.Z-R*R) + 2*R*R*p.X)
	y := -4 * (a*a*b - 2*a*b*p.X + b*b*b - 2*b*b*p.Y + b*(c*c-2*c*p.Z-R*R) + 2*R*R*p.Y)
	z := -4 * c * (a*a - 2*a*p.X - 2*b*p.Y + c*c - 2*c*p.Z + R*R)

	return vector3{x, y, z}
}

func (torus *Torus) GetMaterial() mats.Material {
	return torus.Mat
}
