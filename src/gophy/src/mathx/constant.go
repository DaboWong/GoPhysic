package mathx

import "math"

const (
	PI = math.Pi
)

func RadiusToEuler(val float64) float64 {
	return val / PI * 180
}

func EulerToRadius(val float64) float64 {
	return val / 180 * PI
}
