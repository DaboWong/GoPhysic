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
	_, ok := component.(*Transform)
	if !ok {
		return
	}
	_, ok = component.(*Rigidbody)
	if !ok {
		return
	}
	self.components[component] = component
}

func (self *GameObject) RemoveComponent(component Component) {
	if component.(*Transform) != nil {
		return
	}
	if component.(*Rigidbody) != nil {
		return
	}
	delete(self.components, component)
}

func (self *GameObject) GetComponent(compType reflect.Type) Component {
	return nil
}
