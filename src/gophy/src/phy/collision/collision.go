package collision

import "gophy/src/phy/object"

type Collision interface {
	// entity pointer
	Entity() object.Entity

	// the real shape of the collision component
	Shape() interface{} // Collider
}
