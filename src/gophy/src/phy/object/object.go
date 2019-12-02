package object

import (
	"gophy/src/phy/world"
)

//basic object define
type Object interface {
	//world unique instance id
	ID() int32

	//pointer to world belong to
	World() world.World
}




