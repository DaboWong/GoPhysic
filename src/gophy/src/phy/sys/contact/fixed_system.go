package contact

import (
	"gophy/src/phy/conntact"
	"gophy/src/phy/object"
	"gophy/src/phy/world"
	"reflect"
)

type FixedSystem struct {
}

func (self *FixedSystem) Do(world world.World, ctx world.Context) {

	//TODO: 遍历所有的对象， 找到他是否持有fixed joint 然后根据fixed joint 更新他们的位置
	world.RangeObject(func(obj object.Object) bool {

		entity := obj.(object.Entity)

		tp := reflect.TypeOf(&conntact.FixedJoint{})

		fj := entity.Component(tp).(*conntact.FixedJoint)
		if fj == nil {
			return true
		}

		self.do(entity, fj.Target.(object.Entity), fj.FixedLength)

		return true
	})

}

func (self *FixedSystem) do(from, other object.Entity, fixedLength float32) {

	tf := from.Component(reflect.TypeOf(&object.Transform{})).(*object.Transform)
	to := other.Component(reflect.TypeOf(&object.Transform{})).(*object.Transform)


}
