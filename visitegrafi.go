package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type grafo map[string][]string

func main() {
	var s string
	flag := false
	p := 0.8
	g := gen(p)
	count := 0

	fmt.Print("CAMMINO?")
	path("G", "B", g)

	fmt.Println("GRADO:", degree("A", g))

	fmt.Print("STRINGA PER  INIZIARE CONTEGGIO CCC: ")
	fmt.Scan(&s)
	ccc(g, s)

	fmt.Print("STRINGA PER CC: ")
	fmt.Scan(&s)
	cc(g, s)

	fmt.Print("STRINGA PER SPAN: ")
	fmt.Scan(&s)
	span(g, s)

	fmt.Print("STRINGA PER TWOCOLORS: ")
	twocolors(g, "A", count, flag)

}

func dfs1(g *grafo, v, w string, aux *map[string]bool) {
	var s string
	flag := false
	(*aux)[v] = true
	for _, v2 := range (*g)[v] {
		if (*aux)[v2] != true {
			if v2 == w {
				(*aux)[v2] = true
				flag = true
				fmt.Println("SI")
				return
			}
		}
		s = v2
	}
	if flag == false {
		fmt.Println("NO")
		return
	}
	dfs1(g, s, w, aux)
}
func dfs2(g grafo, v string, aux *map[string]bool, colori *map[string]string, count int, ciclo bool) {

	(*aux)[v] = true

	if count == 0 {
		(*colori)[v] = "nero"
		count++
	}

	for _, v2 := range g[v] {
		if (*aux)[v2] != true {
			count++
			if count%2 == 0 {
				(*colori)[v2] = "nero"
			} else {
				(*colori)[v2] = "bianco"
			}
		}
		for _, v3 := range g[v2] {
			// se il colore del mio vertice è uguale a uno dei suoi adiacenti finisce il ciclo funzione nn clorabile
			fmt.Println(v2, v3)
			fmt.Println((*colori)[v2], (*colori)[v3])
			if (*colori)[v2] == (*colori)[v3] {
				ciclo = true
				break
			}

		}
		if !ciclo {
			dfs2(g, v2, aux, colori, count, ciclo)
		} else {
			break
		}
	}
	fmt.Println(ciclo)
	if ciclo == false {
		fmt.Println("ciclo bicolorabile")
	} else {
		fmt.Println("ciclo non bicolorabile")
	}
}

func bfs1(g *grafo, v string, aux map[string]bool) {
	coda := []string{v}
	tuttiVertici := make([]string, 0)
	aux[v] = true
	var connesse int

	for k := range *g {
		if k != v {
			tuttiVertici = append(tuttiVertici, k)
		}
	}

	for len(tuttiVertici) > 0 {
		for len(coda) > 0 {
			v := coda[0]
			coda = coda[1:]
			fmt.Println(v)
			// contare componenti connesse grafo non orientato
			for _, v2 := range (*g)[v] {
				if !aux[v2] {
					coda = append(coda, v2)
					aux[v2] = true
					for i := 0; i < len(tuttiVertici); i++ {
						if tuttiVertici[i] == v2 {
							tuttiVertici = append(tuttiVertici[:i], tuttiVertici[i+1:]...)
						}
					}
				}
			}
		}
		if len(tuttiVertici) != 0 {
			coda = append(coda, tuttiVertici[0])
		}
		connesse++
	}

	fmt.Println("numero di componenti connesse:", connesse)
}
func bfs2(g *grafo, v string, aux map[string]bool) {
	var cc string
	coda := []string{v}
	aux[v] = true

	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]
		cc = cc + v + " "
		for _, v2 := range (*g)[v] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
	}
	fmt.Println("componenti connesse a", v, "-->", cc)
}

func bfs3(g *grafo, v string, aux map[string]bool) {
	coda := []string{v}
	aux[v] = true
	prec := ""

	for len(coda) > 0 {

		v := coda[0]
		fmt.Println(v)
		coda = coda[1:]
		if prec != "" {
			fmt.Println(prec, "-->", v)
		}
		for _, v2 := range (*g)[v] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
		prec = v
	}
}

func gen(p float64) *grafo {
	grafovertici := creaGrafo()
	file, err := os.Open("dec.txt")
	if err != nil {
		fmt.Println("errore apertura file")
		os.Exit(-1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		prob, _ := strconv.ParseFloat(input[0], 64)
		archi := strings.Split(input[1], "-")
		if prob < p {
			grafovertici[archi[0]] = append(grafovertici[archi[0]], archi[1])
			grafovertici[archi[1]] = append(grafovertici[archi[1]], archi[0])
		}
	}
	return &grafovertici
}
func degree(v string, graph *grafo) int {
	return len((*graph)[v])
}
func creaGrafo() grafo {
	graph := make(map[string][]string)
	return graph
}
func path(v, w string, g *grafo) {
	aux := creaMappa()
	dfs1(g, v, w, aux)
}
func creaMappa() *map[string]bool {
	aux := make(map[string]bool)
	return &aux
}
func ccc(graph *grafo, v string) {
	aux := make(map[string]bool)
	bfs1(graph, v, aux)
}
func cc(graph *grafo, v string) {
	aux := make(map[string]bool)
	bfs2(graph, v, aux)

}
func span(graph *grafo, v string) {
	aux := make(map[string]bool)
	bfs3(graph, v, aux)
}
func twocolors(g *grafo, v string, count int, flag bool) {
	var aux *map[string]bool
	var colori *map[string]string
	auxmap := make(map[string]bool)
	colorimap := make(map[string]string)
	aux = &auxmap
	colori = &colorimap
	dfs2(*g, v, aux, colori, count, flag)

}
