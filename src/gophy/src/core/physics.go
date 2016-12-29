package core

import (
	"gophy/src/mathg"
)

const (
	ForceMode_FORCE int32 = 1
)

var Gravity = mathg.Vector3{
	X: 0,
	Y: -9.81,
	Z: 0,
}
