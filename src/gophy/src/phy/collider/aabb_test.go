package collider

import (
	"gophy/src/mathx/vec3"
	"testing"
)

func TestSplitAABB(t *testing.T) {

	a := AABB3D{
		Min: vec3.Vector3{0, 0, 0},
		Max: vec3.Vector3{100, 100, 100},
	}

	ss := SplitOctAABB(&a)

	for _, v := range ss {
		t.Logf("%+v", v)
	}
}
