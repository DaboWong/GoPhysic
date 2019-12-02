package core

import (

)

type Rigidbody struct {
	*Object
	GameObject      *GameObject
	LinearVelocity  vec3.Vector3
	AngularVelocity vec3.Vector3
	LinearDamp      float32
	AngularDamp     float32
	mass            float32
	inverseMass     float32
	IsKinematic     bool
}

func newRigidbody(gameObject *GameObject) *Rigidbody {
	return &Rigidbody{
		GameObject:     gameObject,
		LinearVelocity: mathg.Vector3_Zero,
		Object:         &Object{},
	}
}

//compoenent interface implements
func (self *Rigidbody) Awake()    {}
func (self *Rigidbody) OnEnable() {}
func (self *Rigidbody) Start()    {}
func (self *Rigidbody) FixedUpdate(deltaTime float32) {
	if self.IsKinematic {
		return
	}

	//速度没有
	if !self.LinearVelocity.Equals(mathg.Vector3_Zero) {
		//translate gameobject position
		//目前： p = v*t
		//TODO 如果加速度加入 : p = v*t + a*t(0.5*a*t^2进行优化）
		self.GameObject.Transform.Position = self.GameObject.Transform.Position.Add(self.LinearVelocity.Multi(deltaTime))
	}
}
func (self *Rigidbody) OnDisable() {}
func (self *Rigidbody) OnDestory() {}
func (self *Rigidbody) GetObject() *Object {
	return self.Object
}
func (self *Rigidbody) AddForce(force mathg.Vector3) {
	self.LinearVelocity.Add(force)
}
func (self *Rigidbody) AddTorque(force mathg.Vector3) {
	self.AngularVelocity.Add(force)
}
func (self *Rigidbody) SetMass(mass float32) {
	self.mass = mass
	self.inverseMass = 1 / mass
}
func (self *Rigidbody) GetMass() float32 {
	return self.mass
}
func (self *Rigidbody) GetInverseMass() float32 {
	return self.inverseMass
}
