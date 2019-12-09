package sys

import "gophy/src/phy/world"

type GravitySystem struct {
}

func (self *GravitySystem) Do(world world.World, ctx world.Context) {

	// TODO: iterator all force generator
	//     apply force to it

}
