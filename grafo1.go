package main

import "fmt"

func main() {
	n := 5
	g := nuovoGrafo(n)

	leggiGrafo(1, 2, g)
	leggiGrafo(0, 2, g)
	leggiGrafo(2, 4, g)
	leggiGrafo(4, 3, g)
	leggiGrafo(3, 1, g)
	leggiGrafo(1, 3, g)
	leggiGrafo(1, 4, g)
	leggiGrafo(1, 3, g)

	fmt.Println(isPresent(1, 3, g))
	fmt.Println(isPresent(4, 1, g))
	stampaGrafo(*g)
}

// -------------------------------------------------------------------------------------------------
// 1.1 SLICE DI LISTE DI ADIACENZA
type grafo struct {
	n         int // numero di vertici
	adiacenti []*linkedList
}
type listNode struct {
	val  int
	next *listNode
}
type linkedList struct {
	head *listNode
}

func creaLista() *linkedList {
	return &linkedList{nil}
}

//restituisce indirizzo di un nuovo grafo con n nodi
func nuovoGrafo(n int) *grafo {
	// array con n-1 nodi
	adiacenti := make([]*linkedList, n)
	for i := 0; i < n; i++ {
		// per ogni vertice ho una linkedlist
		l := creaLista()
		adiacenti[i] = l
	}
	return &grafo{n, adiacenti}
}

// inizializzo memoria per nuovo nodo
func newNode(val int) *listNode {
	return &listNode{val, nil}
}
func addNewNode(l *linkedList, val int) *linkedList {
	p := l.head
	node := newNode(val)
	node.next = p
	l.head = node
	return l
}

//Scrivete una funzione per leggere un grafo da standard input.
//Passo alla funzione n = numero vertici, archi = coppia di vertici
func leggiGrafo(v1, v2 int, g *grafo) {
	// mi restituisce indirizzo del mio grafo di n vertici

	// inserire nuovo nodo nella lista di v1
	g.adiacenti[v1] = addNewNode(g.adiacenti[v1], v2)

}
func stampaGrafo(g grafo) {
	// for su numero di vertici
	for i := 0; i < len(g.adiacenti); i++ {
		fmt.Print("adiacenti del vertice ", i, ":")
		stampaLista(g.adiacenti[i])
		fmt.Println()
	}
}

func stampaLista(l *linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.val, " ")
		p = p.next
	}
}

func searchList(l *linkedList, val int) bool {
	p := l.head
	for p != nil {
		if p.val == val {
			return true
		}
	}
	return false
}

func isPresent(v1, v2 int, g *grafo) bool {
	// prendo la lista di adiacenza di g.adiacenti[v1]
	// scorroper vedere se c'è v2
	for i := 0; i < len(g.adiacenti); i++ {
		if i == v1 {
			if searchList(g.adiacenti[i], v2) {
				return true
			}
		}
	}
	return false
}
