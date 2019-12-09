package force

import (
	"gophy/src/mathx/vec3"
	"gophy/src/phy/object"
)

// 常亮力发生器
// constant force generator
type ConstForceGenerator struct {
	gravity vec3.Vector3
}

// 给物体施加作用力
// apply force to object
func (self *ConstForceGenerator) Apply(obj object.Object) {

}
