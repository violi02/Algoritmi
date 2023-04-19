package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type nodo struct {
	nome string
	arco string
	peso int
}

func main() {
	letturaI()

}

func letturaI() {
	svincoli := make([]nodo, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(-1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		pres1, i := nodoPres(input[0], svincoli)
		pres2, k := nodoPres(input[1], svincoli)
		peso, _ := strconv.Atoi(input[2])
		if pres1 {
			if peso < svincoli[i].peso {
				svincoli[i].peso = peso
				svincoli[i].arco = input[1]
				fmt.Println(peso, svincoli[i].peso)
			}
		} else {
			node := creaNodo(input[0], input[1], peso)
			svincoli = append(svincoli, *node)
		}
		if pres2 {
			fmt.Println(peso, svincoli[i].peso)
			if peso < svincoli[k].peso {
				svincoli[k].peso = peso
				svincoli[k].arco = input[0]
			}
		} else {
			node := creaNodo(input[1], input[0], peso)
			svincoli = append(svincoli, *node)
		}
	}
	for i := 0; i < len(svincoli); i++ {
		fmt.Println(svincoli[i].nome, "-->", svincoli[i].peso, "-->", svincoli[i].arco)
	}

	n_gallerie := cammino("1", "5", svincoli)
	fmt.Println(n_gallerie)

}

func creaNodo(nome, arco string, peso int) *nodo {
	var node nodo
	node.nome = nome
	node.arco = arco
	node.peso = peso
	return &node

}

func nodoPres(val string, s []nodo) (bool, int) {
	for i := 0; i < len(s); i++ {
		if s[i].nome == val {
			return true, i
		}
	}
	return false, 0
}

func cammino(partenza, arrivo string, percorso []nodo) (count int) {

	indice := trovareI(partenza, percorso)

	for i := 0; i < len(percorso); i++ {

		if percorso[indice].nome == arrivo {
			return count
		}
		count++
		indice = trovareI(percorso[indice].arco, percorso)
	}
	count = -1
	return count
}

func trovareI(partenza string, percorso []nodo) int {
	for i := 0; i < len(percorso); i++ {
		if percorso[i].nome == partenza {
			return i

		}
	}
	return -1
}
