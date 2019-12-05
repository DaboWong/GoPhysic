package vec3

import (
	"gophy/src/mathx"
	"math"
)

var (
	Zero  = Vector3{0, 0, 0,}
	Left  = Vector3{1, 0, 0,}
	Right = Vector3{-1, 0, 0,}
	Up    = Vector3{0, 1, 0}
	Down  = Vector3{0, -1, 0}
	One   = Vector3{1, 1, 1}
)

//make a new vector from origin to the end
func Direction(origin, end Vector3) Vector3 {
	return end.Sub(origin).Normalized()
}

//distance between to vector3
func Distance(from, to Vector3) float32 {
	return to.Sub(from).SqrtMagnitude()
}

//get the angle between to vector3
func Angle(from, to Vector3) float32 {
	ll := from.SqrtMagnitude() * to.SqrtMagnitude()
	if ll > 0 {
		dot := from.Dot(to)
		//a dot b equals  |a|*|b|*cos( angle)
		return mathx.RadiusToEuler(float32(math.Acos(float64(dot / ll))))
	}
	return 0
}
