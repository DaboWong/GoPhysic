package core

type Scene struct {
	TimeSetting
	GameObjects []*GameObject
}

func (self *Scene) Run() {

}

func NewScene() *Scene {
	return &Scene{
		//20 frames per second
		TimeSetting.FixedDeltaTime: 0.05,
		TimeSetting.TimeScale:      1,
	}
}
