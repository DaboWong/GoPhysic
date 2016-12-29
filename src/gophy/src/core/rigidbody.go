package core

type Rigidbody struct {
	LinearVelocity  Vector3
	AngularVelocity Vector3
	LinearDamp      float32
	AngularDamp     float32
	Mass            float32
}

func (self *Rigidbody) AddForce(force *Vector3) {
	self.LinearVelocity.Add(force)
}

func (self *Rigidbody) AddTorque(force *Vector3) {
	self.AngularVelocity.Add(force)
}
