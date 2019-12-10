package collider

import "gophy/src/mathx/vec3"

type AABB3D struct {
	Min vec3.Vector3
	Max vec3.Vector3
}

func (self *AABB3D) Size() vec3.Vector3 {
	return vec3.Vector3{
		X: self.Max.X - self.Min.X,
		Y: self.Max.Y - self.Min.Y,
		Z: self.Max.Z - self.Min.Z,
	}
}

// 把一个AABB 划分为8个小的AABB
func SplitOctAABB(aabb *AABB3D) []*AABB3D {
	r := make([]*AABB3D, 0)

	size := vec3.Sub(aabb.Max, aabb.Min)
	size = size.Multi(0.5) // split size

	offset := []vec3.Vector3{
		{X: 0, Y: 0, Z: 0},
		{X: size.X, Y: 0, Z: 0},
		{X: 0, Y: size.Y, Z: 0},
		{X: size.X, Y: size.Y, Z: 0},

		{X: 0, Y: 0, Z: size.Z},
		{X: size.X, Y: 0, Z: size.Z},
		{X: 0, Y: size.Y, Z: size.Z},
		{X: size.X, Y: size.Y, Z: size.Z},
	}

	for i := 0; i < 8; i++ {
		a := &AABB3D{
			Min: vec3.Add(aabb.Min, offset[i]),
		}

		a.Max = vec3.Add(a.Min, size)

		r = append(r, a)
	}
	return r
}
