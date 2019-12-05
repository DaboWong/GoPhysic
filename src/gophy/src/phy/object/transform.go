package object

import (
	"gophy/src/mathx/mat"
	"gophy/src/mathx/quat"
	"gophy/src/mathx/vec3"
)

type Transform struct {
	//local position of the object
	localPosition vec3.Vector3

	//local rotation of the object
	localRotation quat.Quaternion

	//world position
	worldPosition vec3.Vector3

	//world rotation
	worldRotation quat.Quaternion

	//local to parant matrix
	localToParentMatrix mat.Matrix4x4
}

//TODO: world position of the object
func (self*Transform) Position()  vec3.Vector3{
	return self.localPosition
}

//TODO: set world position of the object
func (self*Transform) SetPosition( vector3 vec3.Vector3 )  vec3.Vector3{
	self.localPosition = vector3
}