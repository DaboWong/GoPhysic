package core

import (
	"gophy/src/mathg"
)

type Transform struct {
	GameObject   *GameObject
	Position     mathg.Vector3
	Rotation     mathg.Quaternion
	LocalToWorld mathg.Matrix4x4
	WorldToLocal mathg.Matrix4x4
}

func newTransform(gameObject *GameObject) *Transform {
	return &Transform{
		GameObject: gameObject,
		Position:   mathg.Vector3_Zero,
		Rotation:   mathg.Quaternion_Identity,
	}
}

func (self *Transform) Awake()                        {}
func (self *Transform) OnEnable()                     {}
func (self *Transform) Start()                        {}
func (self *Transform) FixedUpdate(deltaTime float32) {}
func (self *Transform) OnDisable()                    {}
func (self *Transform) OnDestory()                    {}
