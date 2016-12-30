package core

type IForceGenerator interface {
	RegisterRigidbody(rigidbody *Rigidbody)
	UpdateForce(deltaTime float32)
}

//gravity generator
type GravityGenerator struct {
	rigidbody *Rigidbody
}

func (self *GravityGenerator) RegisterRigidbody(rigidbody *Rigidbody) {
	self.rigidbody = rigidbody
}

func (self *GravityGenerator) UpdateForce(deltaTime float32) {
	if self.rigidbody.IsKinematic {
		return
	}
	//G = g * m
	self.rigidbody.AddForce(Gravity.Multi(self.rigidbody.GetMass()))
}

//Linear damper genrator
type DamperGemertator struct {
	rigidbody *Rigidbody
}

func (self *DamperGemertator) RegisterRigidbody(rigidbody *Rigidbody) {
	self.rigidbody = rigidbody
}

func (self *DamperGemertator) UpdateForce(deltaTime float32) {
	//take velocity as force
	force := self.rigidbody.LinearVelocity
	magnitude := force.SqrtMagnitude()
	//F(damp)=-velocity*(k1*|p|+k2*|p|*|p|)
	//为了简化 只取LinearDamp
	dragCoeff := self.rigidbody.LinearDamp*magnitude + self.rigidbody.LinearDamp*magnitude*magnitude

	force.Multi(-dragCoeff)

	self.rigidbody.AddForce(force)
}
