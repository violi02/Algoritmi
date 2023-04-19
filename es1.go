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

const lunghezza = 10

func main() {
	var albero *bitree
	albero = standard()
	fmt.Println("inorder:")
	// passo ad ogni funzion radice dell'albero ( primo elemento )
	inorder(albero.root)
	fmt.Println("preorder:")
	preorder(albero.root)
	fmt.Println("postorder:")
	postorder(albero.root)
}

// generare sequenza di interi e inserirli in una slice
/*func generaSeq(int) []int {
	seqint := make([]int, lunghezza)
	seqint = rand.Perm(10)
	return seqint
}*/

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

// crea Nodo
func newNode(n int) *bitreeNode {
	return &bitreeNode{nil, nil, n}
}

/*func costruisciAlbero(seq []int) []*bitreeNode {
	albero := make([]*bitreeNode, 0)
	for i := 0; i < len(seq); i++ {
		nodo := newNode(seq[i])
		albero = append(albero, nodo)
	}
	return albero
}*/

func inorder(node *bitreeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Println(node.val)
	inorder(node.right)
}

// radice, tutto a sx, poi a dx
func preorder(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	preorder(node.left)
	preorder(node.right)
}

func postorder(node *bitreeNode) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Println(node.val)
}
