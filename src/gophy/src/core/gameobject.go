package core

import (
	"reflect"
)

type GameObject struct {
	Name       string
	components map[Component]Component
	Transform  *Transform
	Rigidbody  *Rigidbody
}

func NewGameObject() *GameObject {
	gameObject := &GameObject{}

	gameObject.AddComponent(NewTransform())
	gameObject.AddComponent(NewRigidbody())

	return gameObject
}

func (self *GameObject) AddComponent(component Component) {

	//you cann't add a rigidbody or transform to gameobject!
	switch (component.(type)); {
	case Rigidbody, Transform:
		return
	}

	self.components[component] = component
}

func (self *GameObject) RemoveComponent(component Component) {

	//alson you can't remove the rigidbody or transform
	switch (component.(type)); {
	case Rigidbody, Transform:
		return
	}

	delete(self.components, component)
}

func (self *GameObject) GetComponent(compType reflect.Type) Component {
	return nil
}
