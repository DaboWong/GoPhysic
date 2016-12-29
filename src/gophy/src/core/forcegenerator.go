package core

type IForceGenerator interface {
	UpdateForce(step float32)
}

type GravityGenerator struct {
	Rigidbody
}

func (self *GravityGenerator) UpdateForce(step float32) {
	self.Rigidbody.AddForce(Gravity.Multi(self.Rigidbody.Mass))
}
