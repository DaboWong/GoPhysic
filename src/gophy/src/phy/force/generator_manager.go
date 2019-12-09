package force

import (
	"gophy/src/phy/world"
)

//force generator manager
type GeneratorManager struct {
	generators []Generator
	world      world.World
}

//遍历所有的物体，施加作用力
// range all generator to simulate all objects
func (self *GeneratorManager) Range() {

}
