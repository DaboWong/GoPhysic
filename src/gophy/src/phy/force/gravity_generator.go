package force

import (
	"gophy/src/phy/object"
)

//常亮力发生器
//constant force generator
type GravityForceGenerator struct {
	*ConstForceGenerator
}

// 给物体施加作用力
// apply force to object
func (self *GravityForceGenerator) Apply(obj object.Object) {
	if o, ok := obj.(GravityObject); ok && o.IsUseGravity() {
		obj.(ForceApplier).AddForce(self.gravity, Linear)
	}
}
