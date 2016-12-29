package core

type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

func QuaternionDot(a, b *Quaternion) float32 {
	return 0
}

func FromRotation(from, to *Vector3) *Quaternion {
	return nil
}

func QuaternionFromEulerAngle(x, y, z float32) *Quaternion {
	return nil
}