package octtree

import (
	"gophy/src/algo"
	"gophy/src/phy/collider"
)

type nodeHandler func(*Node) bool

// AABB bord parse search oct tree
// Oct Tree
type OctTree struct {
	// eight children nodes
	*Node

	// biggest bounding box
	MaxAABB *collider.AABB3D

	// will continue create
	createHandler nodeHandler
}

// create oct tree by max aabb and create handler

func NewOctTree(maxAABB *collider.AABB3D, createHandler nodeHandler) *OctTree {
	o := &OctTree{
		MaxAABB:       maxAABB,
		createHandler: createHandler,
	}

	// create root node
	o.Node = newNode(o, nil, maxAABB)

	// create root's 8 children nodes
	o.CreateChildren(o.Node)

	return o
}

// iterator all nodes

func (self *OctTree) Range(f func(tree algo.Tree, pNode interface{})) {

	// walk root node
	f(self, self.Node)

	// iterate children to walk the tree
	for _, v := range self.Children {

		// range every child node
		v.Range(f)
	}
}

// crud

func (self *OctTree) CreateChildren(parent *Node) {

	// test can create
	if !self.createHandler(parent) {
		return
	}

	// split to 8 aabb
	aabb3D := collider.SplitOctAABB(parent.AABB3D)

	// create ndoe by aabb
	parent.Children = make([]*Node, len(aabb3D))
	for k, v := range aabb3D {
		n := newNode(self, parent, v)
		parent.Children[k] = n
		self.CreateChildren(n)
	}
}
