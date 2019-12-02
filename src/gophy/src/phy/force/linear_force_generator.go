package force

import "gophy/src/phy/object"

type Generator interface {
	Apply(obj object.Object)
}
