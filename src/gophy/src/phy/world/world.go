package world

// world will contain all objects in some collections
// context will contain object id

type World interface {
	Update()
	AddContextGroup(group ContextGroup)
	AddContext(group string, ctx Context)
}

// world process progress:

// 1. update all object by force generator

// 2. simulate all object's movement or rotate

// 3. solve contact

// 4. solve collision

// 5. update all object position and rotation

// struct

// World -> 聚合  -> context
// World -> 聚合 -> system
/*
foreach sys in system =>
	for each group in context group
		system.Do( goup or context )
*/

