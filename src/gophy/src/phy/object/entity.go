package object

import "reflect"

type Entity interface {
	// pointer to basic object
	Object

	// component of the object
	Component(p reflect.Type) interface{}

	// components by name
	ComponentsByName(name interface{}, receiver interface{})

	// get components by type
	ComponentsByType(t interface{}, receiver interface{})

	// add component
	AddComponent(val interface{}) interface{}

	// remove component by component
	RemoveComponent(val interface{})

	// remove component by name
	RemoveComponentByName(name interface{})
}
