package main

import "fmt"

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}
type bitree struct {
	root *bitreeNode
}

func main() {
	var albero *bitree
	albero = standard()
	stampaAlbero(albero.root)
	a := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	radice := arr2tree(a, 0)
	fmt.Println()
	stampaAlbero(radice)

}

func newNode(n int) *bitreeNode {
	return &bitreeNode{nil, nil, n}
}

func standard() *bitree {
	t := &bitree{nil}
	t.root = &bitreeNode{nil, nil, 78}
	t.root.left = newNode(54)
	t.root.left.right = newNode(90)
	t.root.left.right.left = newNode(19)
	t.root.left.right.right = newNode(95)
	t.root.right = newNode(21)
	t.root.right.left = newNode(16)
	t.root.right.left.left = newNode(5)
	t.root.right.right = newNode(19)
	t.root.right.right.left = newNode(56)
	t.root.right.right.right = newNode(43)

	return t
}

func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}
	nodo := newNode(a[i])
	// in a[0] ho il valore della radice del mio albero
	root = nodo
	root.left = arr2tree(a, 2*i+1)
	root.right = arr2tree(a, 2*i+2)

	return root
}

func inorder(node *bitreeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Println(node.val)
	inorder(node.right)
}
