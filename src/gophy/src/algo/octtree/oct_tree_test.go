package octtree

import (
	"fmt"
	"gophy/src/algo"
	"gophy/src/mathx/vec3"
	"gophy/src/phy/collider"
	"testing"
	"time"
)

func createOctTree() *OctTree {
	var size float64 = 100
	var minSize float64 = 50

	var timeStart = time.Now()

	defer func() {
		fmt.Printf("cost time: %+v", time.Now().Sub(timeStart))
	}()

	maxAABB := &collider.AABB3D{
		Min: vec3.Vector3{0, 0, 0},
		Max: vec3.Vector3{size, size, size},
	}

	octTree := NewOctTree(maxAABB, func(node *Node) bool {
		sz := node.AABB3D.Size()
		if sz.X <= minSize || sz.Y <= minSize || sz.Z <= minSize {
			return false
		}
		return true
	})

	return octTree
}

func TestNewOctTree(t *testing.T) {

	var octTree = createOctTree()

	var nodeToString func(node *Node)

	nodeToString = func(node *Node) {
		t.Logf("depth: %v n: %v", node.Depth, node.AABB3D)
		for k, v := range node.Children {
			t.Logf("%v | %+v", k, v)
		}

		for _, v := range node.Children {
			nodeToString(v)
		}
	}

	nodeToString(octTree.Node)
}

func TestOctTree_Walk(t *testing.T) {

	var octTree = createOctTree()

	octTree.Walk(func(tree algo.Tree, pNode interface{}) {

		node := pNode.(*Node)

		t.Logf("%d, %+v", node.Depth, node.AABB3D)
	})

}
