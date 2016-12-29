package core

type IForceGenerator interface {
	UpdateForce(deltaTime float32)
}

type GravityGenerator struct {
	Rigidbody
}

func (self *GravityGenerator) UpdateForce(deltaTime float32) {
	self.Rigidbody.AddForce(Gravity.Multi(self.Rigidbody.Mass))
}
