package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	leggiInput()
}

func leggiInput() {
	file, err := os.Open("sei.txt")
	if err != nil {
		fmt.Println("errore nella lettura del file")
		os.Exit(-1)
	}
	defer file.Close()

	var input string
	var pos_carattere int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// input tutto su una riga, unica stringa
		input = scanner.Text()
		pos_carattere = trovaSequenza(input)
	}

	fmt.Println(pos_carattere)
}

// non gli interessa che stringa,ma solo l'ultima posizione del carattere della stringa
// di 4 caratteri differenti

//nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
func trovaSequenza(input string) int {
	var flag bool
	var pos int
	aux := []string{string(input[0])}
	count := len(aux)
	for i := 0; i < len(input); i++ {
		// posizione in cui si trova ultimo carattere
		pos++
		flag = true
		for j := 0; j < len(aux); j++ {
			if string(input[i]) == aux[j] {
				flag = false
				if j == 0 {
					aux = aux[1:]
					aux = append(aux, string(input[i]))
				} else {
					aux = append(aux, string(input[i]))
					s := aux[j+1:]
					aux = s
				}
				count = len(aux)
				break
			}
		}
		if flag == true {
			aux = append(aux, string(input[i]))
			count = len(aux)
		}
		if count == 14 {
			return pos
		}
	}

	return -1
}
