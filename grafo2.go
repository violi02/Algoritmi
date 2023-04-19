// implementazione di grafi con verttici piu articolati
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	n := 5
	g := graphNew(n)
	leggiGraph(*g)
	leggiGraph2(*g)
	printGrafo(g)
	printHobby(g, "A")
}

type vertice struct {
	nome  string
	eta   string
	hobby []string
}

type graf struct {
	vertici   map[string]*vertice
	adiacenti map[string][]string
}

func graphNew(n int) *graf {
	var g graf
	vertici := make(map[string]*vertice)
	adiacenti := make(map[string][]string)
	g.vertici = vertici
	g.adiacenti = adiacenti

	return &g
}

func leggiGraph(g graf) {

	file, err := os.Open("twitter.txt")
	if err != nil {
		fmt.Println("ERRORE APERTURA FILE")
		os.Exit(-1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var v vertice
		input := strings.Split(scanner.Text(), " ")
		input2 := strings.Split(input[1], ",")
		v.nome = input2[0]
		v.eta = input2[1]
		hobby := make([]string, 0)
		hobby = append(hobby, input2[2:]...)
		v.hobby = hobby
		fmt.Println(v)
		g.vertici[input[0]] = &v
	}

	// lettura informazioni seguaci-seguiti

}

func leggiGraph2(g graf) {
	file2, err := os.Open("info.txt")
	if err != nil {
		fmt.Println("ERRORE APERTURA FILE")
		os.Exit(-1)
	}
	defer file2.Close()
	scanner := bufio.NewScanner(file2)
	for scanner.Scan() {
		info := strings.Split(scanner.Text(), "-")
		g.adiacenti[info[0]] = append(g.adiacenti[info[0]], info[1])
		fmt.Println(g.adiacenti[info[0]])
	}
}

func printGrafo(g *graf) {
	// prima stampo informazioni per ogni vertice
	for k, val := range g.vertici {
		fmt.Println("UTENTE", k)
		fmt.Println("NOME:", val.nome)
		fmt.Println("ETA:", val.eta)
		fmt.Println("HOBBY:", val.hobby)
		fmt.Println("//////////////")
	}

	for k, val := range g.adiacenti {
		fmt.Print("UTENTE ", k, " SEGUE")
		fmt.Println(val)
	}
}

func printHobby(g *graf, id string) {
	fmt.Println("HOBBY DI", id, ":", g.vertici[id].hobby)
	for _, val := range g.adiacenti[id] {
		fmt.Print("HOBBY DI ", val, ":", g.vertici[val].hobby)
	}
}

func seguiti(g *graf, id string) {

}
