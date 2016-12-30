package core

import (
	"log"
	"time"
)

//game world of physics
type World struct {
	TimeSetting
	GameObjects   map[*GameObject]*GameObject
	forceGenMgr   *ForceGeneratorManager
	lastFrameTime int64
	quit          bool
	finChan       chan bool
}

func NewWorld() *World {
	return &World{
		forceGenMgr: NewForceGenManager(),
		TimeSetting: TimeSetting{
			UnScaledDeltaTime: (int64)(time.Second) / 20,
			TimeScale:         1,
		},
		finChan:     make(chan bool, 1),
		GameObjects: make(map[*GameObject]*GameObject),
	}
}

//start up world
func (self *World) Startup() {
	self.TimeStartup = time.Now()
}

//Run the entire world
func (self *World) Run() {
	self.resetTimer()
	for {
		select {
		case <-self.TimeSetting.Timer.C:
			//log.Println("update")
			self.update((float32)((float32)(self.DeltaTime()) / (float32)(time.Second)))
			self.resetTimer()

		case msg := <-self.finChan:
			switch msg {
			case true:
				self.handleFinish()
			}
		}
	}
}

//fin
func (self *World) Finish() {
	if self.quit {
		log.Println("duplicated quit world.")
		return
	}
	self.quit = true
	log.Println("world exist!")
	self.finChan <- true
}

//handle finish, cleanup resources.
func (self *World) handleFinish() {
	log.Println("world is finished! wait for cleanup resources!")

	log.Println("close circle chan!")
	close(self.finChan)
}

func (self *World) update(deltaTime float32) {
	//update force generator
	if self.forceGenMgr != nil {
		self.forceGenMgr.Update(deltaTime)
	}
}

//add gameobject to gameobject map.
func (self *World) addGameObject(gameObject *GameObject) {

	//object already added
	if _, ok := self.GameObjects[gameObject]; ok {
		return
	}

	self.GameObjects[gameObject] = gameObject
}

//create a GameObject and add it to the world
func (self *World) NewGameObject() *GameObject {
	log.Println("create new gameObject")
	gameObject := newGameObject()
	self.addGameObject(gameObject)
	log.Printf("game object count:%v", len(self.GameObjects))
	return gameObject
}
