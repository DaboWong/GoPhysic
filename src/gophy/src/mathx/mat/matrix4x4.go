package mat

type Matrix4x4 struct {
	m00, m01, m02, m03,
	m10, m11, m12, m13,
	m20, m21, m22, m23,
	m30, m31, m32, m33 float32
}

//identity matrix
var (
	IdentityMatrix4x4 Matrix4x4 = Matrix4x4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
)

// 创建一个新的矩阵
func NewMatrix4x4() Matrix4x4 {
	return IdentityMatrix4x4
}

//TODO
func (self *Matrix4x4) Mul(other Matrix4x4) Matrix4x4 {
	return IdentityMatrix4x4
}

//TODO 求逆矩阵
func (self *Matrix4x4) Inverse() Matrix4x4 {
	return IdentityMatrix4x4
}

//TODO 求转置矩阵
func (self *Matrix4x4) Transpose() Matrix4x4 {
	return IdentityMatrix4x4
}
