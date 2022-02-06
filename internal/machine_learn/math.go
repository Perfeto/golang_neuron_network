package machine_learn

import "math"

func sigmoid(x float32) float32 {
	return float32(1 / (math.Exp(float64(x*-1)) + 1))
}

func sigmoidDerivative(errValue float32) float32 {
	return sigmoid(errValue) * (1 - sigmoid(errValue))
}
