package core

type Matrix4x4 struct {
	m00, m01, m02, m03,
	m10, m11, m12, m13,
	m20, m21, m22, m23,
	m30, m31, m32, m33 float32
}

func CreateIndityMatrix4x4() *Matrix4x4 {
	return &Matrix4x4{
		m30: 1, m31: 0, m32: 0, m33: 0,
		m20: 0, m21: 1, m22: 0, m23: 0,
		m10: 0, m11: 0, m12: 1, m13: 0,
		m00: 0, m01: 0, m02: 0, m03: 1,
	}
}

//TODO
func (self *Matrix4x4) Mul(other *Matrix4x4) *Matrix4x4 {
	return nil
}

//TODO
func (self *Matrix4x4) Inverse() *Matrix4x4 {
	return nil
}

//TODO
func (self *Matrix4x4) Transpose() *Matrix4x4 {
	return nil
}
