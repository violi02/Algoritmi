package main

import "fmt"

type pila struct {
	elementi []int
	ultimo   *int
}

func main() {
	var coda []int
	enqueue(&coda, -1, 7)
	enqueue(&coda, 5, 7)
	enqueue(&coda, 13, 7)
	enqueue(&coda, 30, 7)
	enqueue(&coda, 1, 7)
	dequeue(&coda)
	enqueue(&coda, 2, 7)
	enqueue(&coda, 4, 7)
	enqueue(&coda, 9, 7)

	printCoda(coda)

}

//implemento una pila limitata con le sue principali funzioni
// LIFO
func push(pila *[]int, n int, piena int) {
	// se la mia pila non è piena
	if !isFull(pila, piena) {
		*pila = append(*pila, n)
	} else {
		fmt.Println("Non posso aggiungere", n, "perchè la pila è piena")
	}
}
func pop(pila *[]int) {
	if !isEmpty(pila) {
		*pila = (*pila)[:len(*pila)-1]
	} else {
		fmt.Println("La coda è vuota")
	}
}
func isEmpty(pila *[]int) bool {
	if len(*pila) == 0 {
		return true
	}
	return false
}
func isFull(pila *[]int, piena int) bool {
	if len(*pila) == piena {
		return true
	}
	return false
}
func printPila(pila []int) {
	for i := 0; i < len(pila); i++ {
		fmt.Println(pila[i])
	}
}

func enqueue(coda *[]int, n, piena int) {
	// se è piena non aggiungo,,altrimenti aggiungo alla fine
	if !isFull(coda, piena) {
		*coda = append(*coda, n)
	} else {
		fmt.Println("Non posso aggiungere", n, "perchè la coda è piena")
	}
}

func dequeue(coda *[]int) {
	if !isEmpty(coda) {
		*coda = (*coda)[1:]
	} else {
		fmt.Println("La coda è vuota")
	}
}
func printCoda(coda []int) {
	for i := 0; i < len(coda); i++ {
		fmt.Println(coda[i])
	}
}
