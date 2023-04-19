package main

import (
	"fmt"
	"strings"
)

type bitreeNode struct {
	val int
	sx  *bitreeNode
	dx  *bitreeNode
}

type bitree struct {
	root *bitreeNode
}

func main() {
	var albero *bitree
	albero = implementazione()
	stampaAlberoASommario(albero.root, 0)
	stampaPreOrder(albero.root, -1)

}

func implementazione() *bitree {
	// ricordarsi di allocare memoria ogni volta che si crea un nodo da inserire nell'albero
	t := &bitree{nil}
	t.root = &bitreeNode{78, nil, nil}
	t.root.sx = newNode(54)
	t.root.dx = newNode(21)
	t.root.sx.dx = newNode(90)
	t.root.sx.dx.sx = newNode(19)
	t.root.sx.dx.dx = newNode(95)
	t.root.dx.sx = newNode(16)
	t.root.dx.sx.sx = newNode(5)
	t.root.dx.dx = newNode(19)
	t.root.dx.dx.sx = newNode(56)
	t.root.dx.dx.dx = newNode(43)
	return t
}

func newNode(n int) *bitreeNode {
	return &bitreeNode{n, nil, nil}
}

func preorder(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print("*")
	fmt.Print(node.val)
	fmt.Println()
	if node.sx != nil {
		preorder(node.sx)
	} else if node.dx != nil {
		fmt.Println("*")
	}
	if node.dx != nil {
		preorder(node.dx)
	} else if node.sx != nil {
		fmt.Println("*")
	}

}

func stampaAlberoASommario(node *bitreeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println(node.val)
	if node.sx != nil {
		stampaAlberoASommario(node.sx, spaces+1)
	} else if node.dx != nil {
		fmt.Println(strings.Repeat(" ", spaces), "*")
	}
	if node.dx != nil {
		stampaAlberoASommario(node.dx, spaces+1)
	} else if node.sx != nil {
		fmt.Println(strings.Repeat(" ", spaces), "*")
	}
}

// stampaAlbero con visita
func stampaPreOrder(node *bitreeNode, spaces int) {
	spaces = spaces + 1
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	preorder(node)
}
