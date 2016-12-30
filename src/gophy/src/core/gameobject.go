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

func newGameObject() *GameObject {
	gameObject := &GameObject{}
	gameObject.Name = "GameObject"
	gameObject.AddComponent(newTransform(gameObject))
	gameObject.AddComponent(newRigidbody(gameObject))

	return gameObject
}

func (self *GameObject) AddComponent(component Component) {
	//you can't add a rigidbody or transform to a gameobject!
	switch component.(type) {
	case *Rigidbody, *Transform:
		return
	}

	self.components[component] = component
}

func (self *GameObject) RemoveComponent(component Component) {

	//also you can't remove rigidbody or transform from a game object
	switch component.(type) {
	case *Rigidbody, *Transform:
		return
	}

	delete(self.components, component)
}

func (self *GameObject) GetComponent(compType reflect.Type) Component {
	return nil
}
