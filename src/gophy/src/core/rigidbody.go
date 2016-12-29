package core

import (
	"gophy/src/mathg"
)

type Rigidbody struct {
	GameObject      *GameObject
	LinearVelocity  *mathg.Vector3
	AngularVelocity *mathg.Vector3
	LinearDamp      float32
	AngularDamp     float32
	Mass            float32
}

func NewRigidbody() *Rigidbody {
	return &Rigidbody{}
}
func (self *Rigidbody) Awake()    {}
func (self *Rigidbody) OnEnable() {}
func (self *Rigidbody) Start()    {}
func (self *Rigidbody) FixedUpdate(deltaTime float32) {
	self.GameObject.Transform.Position = self.GameObject.Transform.Position.Add(self.LinearVelocity.Multi(deltaTime))
}
func (self *Rigidbody) OnDisable() {}
func (self *Rigidbody) OnDestory() {}

func (self *Rigidbody) AddForce(force *mathg.Vector3) {
	self.LinearVelocity.Add(force)
}
func (self *Rigidbody) AddTorque(force *mathg.Vector3) {
	self.AngularVelocity.Add(force)
}
