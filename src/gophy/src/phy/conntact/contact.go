package conntact

import "gophy/src/phy/object"

type Contact interface {
	//from object
	ObjectA() object.Object

	//to object
	ObjectB() object.Object

	// solve the object
	Solve()
}
