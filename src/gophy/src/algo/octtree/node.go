package octtree

import (
	"gophy/src/algo"
	"gophy/src/phy/collider"
)

type DebugInfo struct {
	Depth int
}

// node for bounding box tree
type Node struct {
	*collider.AABB3D          // axis align bounding box
	Children         []*Node  // all children
	tree             *OctTree // tree  pointer
	parent           *Node    // parent node , nil = root
	DebugInfo
}

func newNode(tree *OctTree, parent *Node, aabb *collider.AABB3D) *Node {
	n := &Node{
		AABB3D:   aabb,
		Children: make([]*Node, 0),
		tree:     tree,
		parent:   parent,
	}

	if parent != nil {
		n.Depth = parent.Depth + 1
	}

	return n
}

func (self *Node) Range(f func(tree algo.Tree, pNode interface{})) {

	f(self.tree, self)

	for _, v := range self.Children {
		v.Range(f)
	}
}
