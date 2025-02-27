package lanchester

import "math"

// SquareLaw simula il modello quadratico di Lanchester
func SquareLaw(R0, B0, rS, bS float64, T, dt int) ([]float64, []float64) {
	R := []float64{R0}
	B := []float64{B0}

	for t := 0; t < T; t += dt {
		R = append(R, R[t]-bS*B[t])
		B = append(B, B[t]-rS*R[t])
		if R[len(R)-1] < 1e-6 || B[len(B)-1] < 1e-6 {
			break
		}
	}
	return R, B
}

// LinearLaw simula il modello lineare di Lanchester
func LinearLaw(R0, B0, rL, bL float64, T, dt int) ([]float64, []float64) {
	R := []float64{R0}
	B := []float64{B0}

	for t := 0; t < T; t += dt {
		R = append(R, R[t]-bL*B[t]*R[t])
		B = append(B, B[t]-rL*B[t]*R[t])
		if R[len(R)-1] < 1e-6 || B[len(B)-1] < 1e-6 {
			break
		}
	}
	return R, B
}

// ModernizedModel simula il modello modernizzato di Lanchester
func ModernizedModel(R0, B0, rL, bL, rS, bS, rF, sR, bF, sB, rI, bI float64, T, dt int) ([]float64, []float64) {
	R := []float64{R0}
	B := []float64{B0}

	for t := 0; t < T; t += dt {
		rValue := R[t] - (1-rF)*sB*rS*B[t]*bI - (1-(1-rF)*sB)*rL*B[t]*R[t]*bI
		bValue := B[t] - (1-bF)*sR*bS*R[t]*rI - (1-(1-bF)*sR)*bL*B[t]*R[t]*rI
		R = append(R, math.Max(0, rValue))
		B = append(B, math.Max(0, bValue))
		if R[len(R)-1] < 1e-6 || B[len(B)-1] < 1e-6 {
			break
		}
	}
	return R, B
}
