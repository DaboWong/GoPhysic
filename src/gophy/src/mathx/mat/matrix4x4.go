package mat

import (
	"gophy/src/mathx/quat"
	"gophy/src/mathx/vec3"
)

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

//TODO 位移旋转缩放
func (self *Matrix4x4) MakeTRS(translate vec3.Vector3, rotation quat.Quaternion, scale vec3.Vector3) {

}

//TODO 位移
func (self *Matrix4x4) Translate(translate vec3.Vector3) {

}

//TOOD 旋转
func (self *Matrix4x4) Rotate(rotation quat.Quaternion) {

}

//TODO 缩放
func (self *Matrix4x4) Scale(scal vec3.Vector3) {

}
