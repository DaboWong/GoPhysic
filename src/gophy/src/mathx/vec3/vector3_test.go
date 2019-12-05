package vec3

import (
	"testing"
)

func TestVector3(t *testing.T) {
	var a = Vector3{1, 0, 0 }
	var b = Vector3{0, 1, 0}

	var dot = a.Dot(b)
	t.Log(dot)

	var angle = Angle(a, b)
	t.Log(angle)

	var cross = a.Cross(b)
	t.Logf("%+v", cross)

	var ma = a.Multi(10)
	t.Logf("%v", ma)

	a.MultiBy(10)
	t.Logf("%v", a)


}
