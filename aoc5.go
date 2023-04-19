package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	letturaInput()

}

func make_m() *map[int][]string {
	stacks := make(map[int][]string)

	//-----CONTENUTO STACK INIZIALE----
	s1 := []string{"J", "H", "G", "M", "Z", "N", "T", "F"}
	s2 := []string{"V", "W", "J"}
	s3 := []string{"G", "V", "L", "J", "B", "T", "H"}
	s4 := []string{"B", "P", "J", "N", "C", "D", "V", "L"}
	s5 := []string{"F", "W", "S", "M", "P", "R", "G"}
	s6 := []string{"G", "H", "C", "F", "B", "N", "V", "M"}
	s7 := []string{"D", "H", "G", "M", "R"}
	s8 := []string{"H", "N", "M", "V", "Z", "D"}
	s9 := []string{"G", "N", "F", "H"}

	//------MAPPA---------
	stacks[1] = s1
	stacks[2] = s2
	stacks[3] = s3
	stacks[4] = s4
	stacks[5] = s5
	stacks[6] = s6
	stacks[7] = s7
	stacks[8] = s8
	stacks[9] = s9

	/*s1 := []string{"Z", "N"}
	s2 := []string{"M", "C", "D"}
	s3 := []string{"P"}

	stacks[1] = s1
	stacks[2] = s2
	stacks[3] = s3*/

	return &stacks
}

func letturaInput() {
	file, err := os.Open("cinque.txt")
	if err != nil {
		fmt.Println("errore nella lettura del file")
		os.Exit(-1)
	}
	defer file.Close()

	stacks := make_m()

	scanner := bufio.NewScanner(file)
	numeri := make([]int, 0)
	for scanner.Scan() {

		// ho tre numeri, li leggo e li inserisco in una slice
		input := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(input); i++ {
			if input[i] >= "0" && input[i] <= "9" {
				n, _ := strconv.Atoi(input[i])
				numeri = append(numeri, n)

			}
		}
		//stacks = move(numeri[0], numeri[1], numeri[2], stacks)
		stacks = moveInorder(numeri[0], numeri[1], numeri[2], stacks)
		numeri = nil
	}

	var message string
	for i := 1; i <= 9; i++ {
		for k := range *stacks {
			if i == k {
				message += (*stacks)[k][len((*stacks)[k])-1]
			}
		}
	}
	fmt.Println(message)

}

func move(numero_crateri, from, to int, m *map[int][]string) *map[int][]string {
	aux := make([]string, 0)
	//j := 1
	for k := range *m {
		if k == from {
			//for len((*m)[from]) > 0 && i != numero_crateri {
			for i := 0; i < numero_crateri; i++ {
				aux = append(aux, (*m)[from][len((*m)[from])-1])
				s := (*m)[from][:len((*m)[from])-1]
				(*m)[k] = s
			}
			(*m)[to] = append((*m)[to], aux...)
		}
	}
	return m
}

func moveInorder(numero_crateri, from, to int, m *map[int][]string) *map[int][]string {
	aux := make([]string, 0)
	for k := range *m {
		if k == from {
			// se il numero di crateri da spostare è uguale alla lunghezzza della stringa
			// sposto tutto direttamente
			if numero_crateri == len((*m)[from]) {
				aux = append(aux, (*m)[from]...)
				(*m)[k] = nil
			}
			if numero_crateri < len((*m)[from]) {
				// elementi che devo togliere
				l := len((*m)[from]) - numero_crateri
				// gli aggiungo ad aux
				aux = append(aux, (*m)[from][l:]...)
				s := (*m)[from][:l]
				(*m)[k] = s
			}

		}
		(*m)[to] = append((*m)[to], aux...)
		aux = nil
	}

	return m
}
