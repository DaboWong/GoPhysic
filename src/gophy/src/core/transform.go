package core

import (
	"gophy/src/mathx"
)

type Transform struct {
	*Object
	GameObject   *GameObject
	Position     vec3.Vector3
	Rotation     quaternion.Quaternion
	LocalToWorld mat.Matrix4x4
	WorldToLocal mathx.Matrix4x4
	Parent       *Transform
}

func newTransform(gameObject *GameObject) *Transform {
	return &Transform{
		GameObject: gameObject,
		Position:   mathx.Vector3_Zero,
		Rotation:   mathx.Quaternion_Identity,
		Object:     &Object{},
	}
}

func (self *Transform) setStart() {
	self.started = true
}

//compoenent interface implements
func (self *Transform) Awake()    {}
func (self *Transform) OnEnable() {}
func (self *Transform) Start() {
}
func (self *Transform) FixedUpdate(deltaTime float32) {}
func (self *Transform) OnDisable()                    {}
func (self *Transform) OnDestory()                    {}
func (self *Transform) GetObject() *Object {
	return self.Object
}
