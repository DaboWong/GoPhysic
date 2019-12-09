package algo

type nodeActionHandler func(tree Tree, node interface{})

// Oct Tree
type OctTree struct {
	// eight children nodes
	nodes []interface{}

	// trigger when node add
	addHandler nodeActionHandler

	// trigger when node remove
	removeHandler nodeActionHandler
}

// iterator all nodes

func (self *OctTree) Walk(f func(tree Tree, pNode interface{}) bool) {
	for _, v := range self.nodes {
		if !f(self, v) {
			break
		}

		// deep sort

		if t, ok := v.(Tree); ok {
			t.Walk(f)
		}
	}
}

// crud

func (self *OctTree) AddNode(node interface{}) {

}

func (self *OctTree) RemoveNode() {

}
