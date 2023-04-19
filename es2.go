package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertice struct {
	nome  string
	eta   int
	hobby []string
}

type grafo struct {
	vertici   map[string]*vertice
	adiacenti map[*vertice][]*vertice
}

func main() {
	leggereGrafo()
}

func leggereGrafo() {
	var v vertice
	graph := creaGrafo()
	// per leggere il grafo devo sia leggere il contenuto del vertice
	// sia i rispettivi collegamenti (archi, follower)
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("errore nell'apertura del file")
		os.Exit(-1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valori := strings.Split(scanner.Text(), "-")
		v.nome = valori[0]
		v.eta, _ = strconv.Atoi(valori[1])
		hobby := strings.Split(valori[2], ",")
		v.hobby = hobby
		aggiungiVert(graph, v)

	}
	//lettura collegamenti
	file2, err := os.Open("collegamenti.txt")
	if err != nil {
		fmt.Println("errore nell'apertura del file")
		os.Exit(-1)
	}
	defer file2.Close()
	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		archi := strings.Split(scanner.Text(), "-")
		aggiungiAdiacenti(graph, archi[0], archi[1])
	}
	stampaGrafo(graph)
	fmt.Println("STAMPA HOBBY SEGUE")
	for k, _ := range graph.vertici {
		stampaHobbySEGUE(k, graph)
	}
}

func aggiungiVert(graph *grafo, v vertice) {
	graph.vertici[v.nome] = &v
}

func aggiungiAdiacenti(graph *grafo, s string, seguace string) {

	for k, v := range graph.vertici {
		// se trovo il valore della stringa corrispondente
		if k == s {
			graph.adiacenti[v] = append(graph.adiacenti[v], graph.vertici[seguace])
		}
		//seguaci = nil
	}
}

func stampaGrafo(graph *grafo) {
	// prima stampo le info di ogni utente
	for k, v := range graph.vertici {
		fmt.Println(k, v.eta, v.hobby)
		for i := 0; i < len(graph.adiacenti[v]); i++ {

			fmt.Println(k, "segue", graph.adiacenti[v][i].nome)
		}
	}
}

func creaGrafo() *grafo {
	var graph grafo
	map1 := make(map[string]*vertice)
	map2 := make(map[*vertice][]*vertice)
	graph.vertici = map1
	graph.adiacenti = map2
	return &graph
}

func stampaHobbyFOLLOWING(s string, graph *grafo) {
	for k, v := range graph.vertici {
		if k == s {
			fmt.Println("hobby di", k, ":", v.hobby)
			for i := 0; i < len(graph.adiacenti[v]); i++ {
				fmt.Println("hobby di chi segue", k, "-->", graph.adiacenti[v][i].nome, ":", graph.adiacenti[v][i].hobby)
			}
		}
	}
}

func stampaHobbyFOLLOWER(s string, graph *grafo) {
	for k, v := range graph.vertici {
		if k == s {
			fmt.Println("hobby di", k, ":", v.hobby)
		}

	}
}
