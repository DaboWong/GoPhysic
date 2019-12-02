package quat

type Quaternion struct {
	X, Y, Z, W float32
}

var (
	Identity = Quaternion{
		X: 0,
		Y: 0,
		Z: 0,
		W: 1,
	}
)
