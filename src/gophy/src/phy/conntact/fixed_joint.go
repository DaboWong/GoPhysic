package conntact

import (
	"gophy/src/phy/object"
)

type FixedJoint struct {
	//entity belong to
	object.Entity

	//other object
	Target object.Object

	//length between two objects
	FixedLength float32

}
