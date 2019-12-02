package quat

import "gophy/src/mathx/vec3"

func Angular(a, b Quaternion) float32 {
	return 0
}

func FromRotation(from, to vec3.Vector3) Quaternion {
	return Identity
}

func QuaternionFromEulerAngle(x, y, z float32) Quaternion {
	return Identity
}
