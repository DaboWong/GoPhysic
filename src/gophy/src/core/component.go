package core

type Component interface {
	Awake()
	OnEnable()
	Start()
	FixedUpdate(deltaTime float32)
	OnDisable()
	OnDestory()
}
