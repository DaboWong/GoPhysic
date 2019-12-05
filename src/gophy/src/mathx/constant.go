package mathx

import "math"

const (
	PI = math.Pi
)

func RadiusToEuler(val float32) float32 {
	return val / PI * 180
}

func EulerToRadius(val float32) float32 {
	return val / 180 * PI
}
