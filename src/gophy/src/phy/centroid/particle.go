package centroid

import (
	"gophy/src/mathx/vec3"
	"gophy/src/phy/world"
)

//实现牛顿定律带质量的粒子
//centroid particle which simulate Newton law of linear and angular
type Particle struct {
	//world belong to
	world world.World

	// instance id
	id int32

	//local position of particle in world
	localPosition vec3.Vector3

	//linear velocity of the particle
	linearVelocity vec3.Vector3

	//linear accelerate of the particle
	linearAccelerate vec3.Vector3

	//linear force of the particle
	linearForce vec3.Vector3

	//angular force of the particle
	angularForce vec3.Vector3

	//mass of the particle
	mass float32

	//use gravity
	userGravity bool

	//impulse of the particle
	impulse vec3.Vector3
}

//实例Id
//instance id
func (self *Particle) ID() int32 {
	return self.id
}

//归属那个世界
//world belong to
func (self *Particle) World() world.World {
	return self.world
}

//物体是否使用重力
//is particle use gravity
func (self *Particle) IsUseGravity() bool {
	return self.userGravity
}

//get the object's mass
func (self*Particle) Mass() float32{
	return self.mass
}