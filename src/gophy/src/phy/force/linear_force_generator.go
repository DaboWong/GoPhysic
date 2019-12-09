package force

import "gophy/src/phy/object"

// force generator
type Generator interface {
	Apply(obj object.Object)
}
