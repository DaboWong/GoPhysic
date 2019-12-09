package algo

// abstract data struct of tree
type Tree interface {
	Walk(func(tree Tree, pNode interface{}) bool)
}
