package collider

import "gophy/src/mathx/vec3"

// Collider base struct
type Collider struct{
	center vec3.Vector3
	size   vec3.Vector3
}

func (self*Collider) Center() vec3.Vector3 {
	return self.center
}

func (self*Collider) SetCenter( c vec3.Vector3 ){
	self.center = c
}

func (self*Collider) Size() vec3.Vector3 {
	return self.size
}

func (self*Collider) SetSize( s vec3.Vector3) {
	self.size = s
}