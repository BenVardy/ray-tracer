package core

import "math"

// SolveGauss solves a system of linear equations
func SolveGauss(m [][]float64, v []float64) []float64 {
	// setup array
	n := len(m)

	a := make([][]float64, n)
	for i, mi := range m {
		a[i] = make([]float64, n)
		copy(a[i], mi)
	}

	b := make([]float64, len(v))
	copy(b, v)

	for i := 0; i < n-1; i++ {

		// Swap rows to ensure that no values are 0 at i, i
		iMax := i
		max := math.Abs(a[i][i])
		for j := i + 1; j < n; j++ {
			if abs := math.Abs(a[j][i]); abs > max {
				iMax, max = j, abs
			}
		}

		if iMax != i {
			a[i], a[iMax] = a[iMax], a[i]
			b[i], b[iMax] = b[iMax], b[i]
		}

		for k := i + 1; k < n; k++ {
			if a[k][i] != 0 {
				factor := a[k][i] / a[i][i]
				for l := range a[k] {
					a[k][l] -= factor * a[i][l]
				}
				b[k] -= factor * b[i]
			}
		}
	}

	for i := len(b) - 1; i >= 0; i-- {
		totalDot := 0.0
		for j := i + 1; j < len(b); j++ {
			totalDot += a[i][j] * b[j]
		}

		b[i] = (b[i] - totalDot) / a[i][i]
	}

	return b
}
