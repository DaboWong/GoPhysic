package conntact

import (
	"gophy/src/mathx/vec3"
	"gophy/src/phy/object"
)

type FixedJoint struct {
	//entity belong to
	object.Entity

	//other object
	Target object.Object

	//length between two objects
	FixedLength float32

	// where to fix origin entity to other object
	// when target is nil, fixpoint will default eqaul to vecto3.zero
	FixPoint vec3.Vector3
}
