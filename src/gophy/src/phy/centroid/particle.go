package centroid

import (
	"gophy/src/mathx/vec3"
	"gophy/src/phy/object"
)

// 实现牛顿定律带质量的粒子
// centroid particle which simulate Newton law of linear and angular
type Particle struct {
	// local position of particle in world
	localPosition vec3.Vector3

	// linear velocity of the particle
	linearVelocity vec3.Vector3

	// linear accelerate of the particle
	linearAccelerate vec3.Vector3

	// linear force of the particle
	linearForce vec3.Vector3

	// angular force of the particle
	angularForce vec3.Vector3

	//m ass of the particle
	mass float32

	// use gravity
	userGravity bool

	// impulse of the particle
	impulse vec3.Vector3

	// entity pointer
	entity object.Entity
}


// 物体是否使用重力
// is particle use gravity
func (self *Particle) IsUseGravity() bool {
	return self.userGravity
}

// get the object's mass
func (self*Particle) Mass() float32{
	return self.mass
}

// entity pointer
func (self*Particle) Entity() object.Entity{
	return self.entity
}