package core

import (
	"gophy/src/mathg"
)

type Transform struct {
	Position     *mathg.Vector3
	Rotation     *mathg.Quaternion
	LocalToWorld *mathg.Matrix4x4
	WorldToLocal *mathg.Matrix4x4
	Parent       *Transform
	GameObject   *GameObject
}

func NewTransform() *Transform {
	return &Transform{}
}

func (self *Transform) Awake()                        {}
func (self *Transform) OnEnable()                     {}
func (self *Transform) Start()                        {}
func (self *Transform) FixedUpdate(deltaTime float32) {}
func (self *Transform) OnDisable()                    {}
func (self *Transform) OnDestory()                    {}
