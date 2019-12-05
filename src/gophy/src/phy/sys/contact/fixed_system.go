package contact

import (
	"gophy/src/mathx/vec3"
	"gophy/src/phy/centroid"
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

		self.do(entity, fj.Target.(object.Entity), fj)

		return true
	})

}

func (self *FixedSystem) do(from, other object.Entity, joint *conntact.FixedJoint) {

	var tf, to *object.Transform

	tf = from.Component(reflect.TypeOf(&object.Transform{})).(*object.Transform)
	if other != nil {
		to = other.Component(reflect.TypeOf(&object.Transform{})).(*object.Transform)
	}
	var fixPoint vec3.Vector3
	if to != nil {
		fixPoint = to.Position()
	}
	fixPoint = joint.FixPoint

	distance := vec3.Distance(tf.Position(), fixPoint)
	if distance > joint.FixedLength {
		var pa, pb *centroid.Particle
		//control objects movement base on the distance between them
		pa = from.Component(reflect.TypeOf(&centroid.Particle{})).(*centroid.Particle)
		if other != nil {
			pb = from.Component(reflect.TypeOf(&centroid.Particle{})).(*centroid.Particle)
		}

		var totalMass = pa.Mass()
		if pb != nil {
			totalMass += pb.Mass()
		}

		//offset distance
		var offset = distance - joint.FixedLength

		//make offset
		var dir = vec3.Direction(tf.Position(), fixPoint)

		//only src will move
		if to == nil {
			offset := dir.Multi(offset)
			tf.SetPosition( tf.Position().Add(offset) )
		}else {
			//TODO: calculate mv
		}
	}
}
