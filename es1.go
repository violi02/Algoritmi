package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type listNode struct {
	val  int
	next *listNode
}

type linkedList struct {
	head *listNode
}

type grafo struct {
	n         int
	adiacenti []*linkedList
}

func main() {
	var n int
	// numero di vertici
	n = 5
	graph := letturaGrafo(n)
	stampaGrafo(graph, n)
	boolArco(4, 1, graph)
	boolArco(2, 3, graph)

}

// creazione di un nuovo grafo
func nuovoGrafo(n int) *grafo {
	var grafo grafo
	s := make([]*linkedList, n)
	grafo.adiacenti = s
	grafo.n = n
	return &grafo
}

func letturaGrafo(n int) *grafo {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("errore nell'apertura del file")
		os.Exit(-1)
	}
	defer file.Close()
	grafo := nuovoGrafo(n)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arco := strings.Split(scanner.Text(), "-")
		primo, _ := strconv.Atoi(arco[0])
		secondo, _ := strconv.Atoi(arco[1])
		// creo la testa se non c'è per quella linkedList
		if grafo.adiacenti[primo] == nil {
			grafo.adiacenti[primo] = &linkedList{nil}
		}
		// devo far puntare la testa al primo nodo della mia linkedList
		// se il suo puntatore è nil
		adiacente := newNode(secondo)
		if grafo.adiacenti[primo].head == nil {
			grafo.adiacenti[primo].head = adiacente
			//altrimenti il mio nuovo nodo punterà al nodo a cui puntava la testa precedentemente
			// e la mia testa punterà al nuovo nodo inserito nella linkedList
		} else {
			adiacente.next = grafo.adiacenti[primo].head
			grafo.adiacenti[primo].head = adiacente
		}
	}
	return grafo
}

func newNode(n int) *listNode {
	var nodo listNode
	nodo.val = n
	nodo.next = nil
	return &nodo
}

func stampaGrafo(graph *grafo, n int) {

	for i := 0; i < n; i++ {
		fmt.Print(i, " : ")
		puntatore := graph.adiacenti[i].head
		for puntatore != nil {
			// valore di ogni mio nodo che ho nella mia linkedlist
			fmt.Print("->", puntatore.val)
			puntatore = puntatore.next
		}
		fmt.Println()
	}
}
func boolArco(x, y int, grafo *grafo) {
	var flag = false
	vertice := grafo.adiacenti[x]

	// inizialmente puntatore è il puntatore alla testa della linkedList del corrispondente vertice
	// che in questo caso specifico è x
	puntatore := vertice.head
	for puntatore != nil {
		if puntatore.val == y {
			flag = true
			fmt.Println("c'è un arco tra", x, "e", y)
		}
		// il mio puntatore inseguito punterà all'elemento successivo presente nella mia linkedlist
		puntatore = puntatore.next
	}
	if flag == false {
		fmt.Println("non c'è un arco tra", x, "e", y)
	}
}
