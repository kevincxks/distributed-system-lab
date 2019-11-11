package binaryTree


type Comparable func(c1 interface{}, c2 interface{}) bool

type BinaryTree struct {
	Node interface{}
	left *BinaryTree
	right *BinaryTree
	lessFun Comparable
}


func New(compareFun Comparable) *BinaryTree {
	return &BinaryTree{
		Node:    nil,
		lessFun: compareFun,
	}
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.Node == nil {
		return nil
	}
	if tree.Node == value {
		return tree
	}
	if tree.lessFun(value, tree.Node) {
		return tree.left.Search(value)
	} else {
		return tree.right.Search(value)
	}
}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.Node == nil {
		tree.Node = value
		tree.left = New(tree.lessFun)
		tree.right = New(tree.lessFun)
		return
	}
	if tree.lessFun(value, tree.Node) {
		tree.left.Insert(value)
	} else {
		tree.right.Insert(value)
	}
}

func (tree *BinaryTree) Max() interface{} {
	if tree.Node == nil || tree.right.Node == nil {
		return tree.Node
	}

	return tree.right.Max()

}
