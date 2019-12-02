package mat

type Matrix3x3 struct {
	m00, m01, m02,
	m10, m11, m12,
	m20, m21, m22 float32
}

var Indentity = &Matrix3x3{
	m00: 1, m01: 0, m02: 0,
	m10: 0, m11: 1, m12: 0,
	m20: 0, m21: 0, m22: 1,
}

func NewMatrix3x3() *Matrix3x3 {
	return &Matrix3x3{
		m00: 1, m01: 0, m02: 0,
		m10: 0, m11: 1, m12: 0,
		m20: 0, m21: 0, m22: 1,
	}
}

func (self *Matrix3x3) MakeTRS(scale vec3.Vector3, rotate quaternion.Quaternion, scale vec3.Vector3) {

}

func (self *Matrix3x3) MakeIndentity() {

}
