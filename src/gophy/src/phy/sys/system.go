package sys

import "gophy/src/phy/world"

//system do
type System interface {
	Do(world world.World, ctx world.Context)
}

type GravitySystem struct {
}

func (self *GravitySystem) Do(world world.World, ctx world.Context) {

	//TODO: iterator all force generator
	//     apply force to it

}
