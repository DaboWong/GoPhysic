package core

import (
	"log"
	"reflect"
)

type GameObject struct {
	*Object
	Name       string
	components map[Component]Component
	Transform  *Transform
	Rigidbody  *Rigidbody
}

//create function
func newGameObject() *GameObject {
	gameObject := &GameObject{
		components: make(map[Component]Component, 0),
		Name:       "GameObject",
		Object:     &Object{},
	}
	gameObject.Transform = gameObject.addComponent(newTransform(gameObject)).(*Transform)
	gameObject.Transform.OnEnable()
	gameObject.Rigidbody = gameObject.addComponent(newRigidbody(gameObject)).(*Rigidbody)
	gameObject.Rigidbody.OnEnable()
	gameObject.SetEnable(true)
	return gameObject
}

//internal add component to gameobject
func (self *GameObject) addComponent(component Component) Component {
	log.Printf("add and awake component %v", reflect.TypeOf(component).String())
	component.Awake()
	self.components[component] = component
	return component
}

//internal remove component from gameobject
func (self *GameObject) removeComponent(component Component) {
	if _, ok := self.components[component]; !ok {
		log.Printf("try remove missing component:%v", reflect.TypeOf(component).String())
		return
	}
	log.Printf("remove and destroy component %v", reflect.TypeOf(component).String())
	component.OnDestory()
	component.OnEnable()
	delete(self.components, component)
}

func (self *GameObject) AddComponent(component Component) Component {
	//you can't add a rigidbody or transform to a gameobject!
	switch component.(type) {
	case *Rigidbody:
		return self.Rigidbody
	case *Transform:
		return self.Transform
	}
	return self.addComponent(component)
}

func (self *GameObject) RemoveComponent(component Component) {
	//also you can't remove rigidbody or transform from a game object
	switch component.(type) {
	case *Rigidbody, *Transform:
		return
	}
	self.removeComponent(component)
}

func (self *GameObject) GetComponent(compType reflect.Type) Component {
	return nil
}
func (self *GameObject) GetObject() *Object {
	return self.Object
}
func (self *GameObject) fixedUpdate(deltaTime float32) {
	for _, v := range self.components {
		if !v.GetObject().isStarted() {
			v.GetObject().setStart()
			v.Start()
			log.Printf("start component:%v", reflect.TypeOf(v).String())
		}
		if v.GetObject().IsEnabled() {
			v.FixedUpdate(deltaTime)
		}
	}
}
