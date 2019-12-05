package world

import "gophy/src/phy/object"

type Context interface {
	//unique name of the context
	Name() string

	//iterator all objects
	RangeObject(f func(object object.Object) bool)

	//find object by id
	ObjectById(id int32) object.Object

	//add new object and will allocate unique intance id
	AddObject(obj object.Object)

	//remove object by id
	RemoveById(id int32)

	//remove object
	Remove(obj object.Object)

	//world belong to
	World() World
}

type ContextGroup interface {
	//add context
	AddContext(ctx Context)

	//find context by name
	Context(name string) Context

	//remove context by name
	RemoveContext(name string)

	//remove all context
	RemoveAllContext()

	//remove context by condition
	RemoveContextBy(f func(ctx Context))
}
