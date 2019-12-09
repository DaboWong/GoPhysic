package sys

import "gophy/src/phy/world"

//system do
type System interface {
	Do(world world.World, ctx world.Context)
}


