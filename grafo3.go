// PROPRIETA' DI UN GRAFO
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	aux := make(map[string]int)
	g := letturaInput()
	grado := degree("A", *g)
	fmt.Println(grado)
	path("C", "D", *g, aux)
	path("A", "F", *g, aux)
}

type g map[string][]string

func gen(p float64) g {
	var grafo g
	vertici := make(map[string][]string)
	grafo = vertici
	return grafo
}

func letturaInput() *g {

	const prob = 0.3
	grafo := gen(prob)
	file, err := os.Open("prob.txt")
	if err != nil {
		fmt.Println("ERRORE NELL'APERTURA DEL FILE")
		os.Exit(-1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		num, _ := strconv.ParseFloat(input[0], 64)
		if num >= prob {
			// se la probabilità è adeguata aggiungo l'arco al grafo
			archi := strings.Split(input[1], "-")
			v1 := archi[0]
			v2 := archi[1]
			grafo[v1] = append(grafo[v1], v2)
			//(*grafo)[archi[0]] = append((*grafo)[archi[0]], archi[1])
		}
	}
	for k, val := range grafo {
		fmt.Println(k, val)
	}
	return &grafo
}

func degree(v string, grafo g) int {
	count := 0
	for i := 0; i < len(grafo[v]); i++ {
		count++
	}
	return count
}

func path(v, w string, grafo g, aux map[string]int) {
	s := grafo[v]
	for i := 0; i < len(s); i++ {
		// se adiacente non ancora visitato
		if aux[s[i]] < 1 {
			aux[s[i]] = 1
			if s[i] == w {
				fmt.Println("esiste un cammino SEMPLICE")
				return

			}
		}
		if aux[s[i]] >= 1 {
			fmt.Println("non esiste un cammino SEMPLICE")
			return
		}
	}
}
func span(v string, aux map[string]bool, grafo g) {
	fmt.Println(v)
	// vertice v visitato
	aux[v] = true
	// scorro gli adiacenti di v
	for _, v2 := range (grafo)[v] {
		// se non ho ancora visitato il nodo richiamo la funzione e lo visito
		if aux[v2] != true {
			span(v2, aux, grafo)
		}
	}
}

//NUMERO COMPONENTI CONNESSE DI UN GRAFO NON ORIENTATO G
func ccc(grafo g) int {
	var cc int
	// slice di vertici
	vertici := make([]string, 0)
	for k, _ := range grafo {
		vertici = append(vertici, k)
	}
	// tengo conto dei vertici visitati
	i := 0
	aux := make(map[string]bool)
	coda := []string{vertici[i]}
	aux[vertici[i]] = true

	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]

		for _, v2 := range grafo[v] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true

				for k := 0; k < len(vertici); k++ {
					if vertici[k] == v2 {
						// devo togliere da vertici elemento in pos k
						v1 := vertici[k+1:]
						vertici = vertici[:k]
						vertici = append(vertici, v1...)
					}
				}
			}
		}
	}
	if len(coda) == 0 && len(vertici) == 0 {
		cc++
	}
	return cc
}

func twocolor(grafo g) bool {

}
