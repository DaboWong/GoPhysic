package algo

import "gophy/src/phy/collider"

// axis align bounding box
type AABBNode struct{
	index int32
	*collider.Cube
}

// AABB Oct tree
type OctTree struct {
	// eight children nodes
	nodes []*AABBNode
}

func(self*OctTree) Walk( f func(tree Tree, pNode interface{}) bool) {
	for _,v :=range self.nodes {
		if !f (self, v ){
			break
		}
	}
}