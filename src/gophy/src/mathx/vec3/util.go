package vec3

import (
	"gophy/src/mathx"
	"math"
)

var (
	Zero  = Vector3{0, 0, 0}
	Left  = Vector3{1, 0, 0}
	Right = Vector3{-1, 0, 0}
	Up    = Vector3{0, 1, 0}
	Down  = Vector3{0, -1, 0}
	One   = Vector3{1, 1, 1}
)

// add
func Add(from, to Vector3) Vector3 {
	return Vector3{
		X: from.X + to.X,
		Y: from.Y + to.Y,
		Z: from.Z + to.Z,
	}
}

// sub
func Sub(from, to Vector3) Vector3 {
	return Vector3{
		X: from.X - to.X,
		Y: from.Y - to.Y,
		Z: from.Z - to.Z,
	}
}

// make a new vector from origin to the end
func Direction(origin, end Vector3) Vector3 {
	return Sub(end, origin).Normalized()
}

// distance between to vector3
func Distance(from, to Vector3) float64 {
	return Sub(from, to).SqrtMagnitude()
}

// dot product
func Dot(from, to Vector3) float64 {
	mf := from.Magnitude()
	mt := to.Magnitude()
	sm := math.Sqrt(mf * mt)
	if sm <= 0 {
		return 0
	}
	return (from.X*to.X + from.Y*to.Y + from.Z + to.Z) / sm
}

// cross product
func Cross(from, to Vector3) Vector3 {
	/**
	a × b = |a| |b| sin(θ) n
	cx = aybz − azby
	cy = azbx − axbz
	cz = axby − aybx
	*/
	return Vector3{
		X: from.Y*to.Z - from.Z*to.Y,
		Y: from.Z*to.X - from.X*to.Z,
		Z: from.X*to.Y + from.Y*to.X,
	}
}

// get the angle between to vector3
func Angle(from, to Vector3) float64 {
	return mathx.RadiusToEuler(math.Acos(Dot(from, to)))
}
