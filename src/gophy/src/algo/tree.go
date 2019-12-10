package algo

// abstract data struct of tree
type Tree interface {
	Range(func(tree Tree, pNode interface{}))
}
