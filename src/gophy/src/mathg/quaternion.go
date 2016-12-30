package mathg

type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

var (
	Quaternion_Identity = Quaternion{
		X: 0,
		Y: 0,
		Z: 0,
		W: 1,
	}
)

func QuaternionDot(a, b Quaternion) float32 {
	return 0
}

func FromRotation(from, to Vector3) Quaternion {
	return Quaternion_Identity
}

func QuaternionFromEulerAngle(x, y, z float32) Quaternion {
	return Quaternion_Identity
}
