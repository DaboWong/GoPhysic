package vec3

import (
	"testing"
)

func TestVector3(t *testing.T) {
	var a = Vector3{1, 0, 0}
	var b = Vector3{0, 2, 0}

	t.Logf("cross: %v", Cross(a, b))
	t.Logf("normalized: %v", b.Normalized())
	b.Normalize()
	t.Logf("normalize: %v", b)
}

func TestBenchMarkFunc(t *testing.T) {
	r := testing.Benchmark(func(b *testing.B) {

	})

	t.Logf("%+v", r)
}
