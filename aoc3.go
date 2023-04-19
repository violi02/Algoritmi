package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	total := letturaInput()
	fmt.Println(total)

}

func letturaInput() int {
	file, err := os.Open("terzo.txt")
	if err != nil {
		fmt.Println("errore nella lettura del file")
		os.Exit(-1)
	}
	defer file.Close()

	var t int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		l := len(input)
		primo := input[:l/2]
		secondo := input[l/2:]
		t = compare2(primo, secondo)
	}

	return t
}

func compare2(primo, secondo string) int {
	total := 0
	var p int
	for i := 0; i < len(primo); i++ {
		for k := 0; k < len(secondo); k++ {
			if primo[i] == secondo[k] {
				c := rune(primo[i])
				// inizio a dare la priorità
				if c >= 'a' && c <= 'z' {
					var base rune = 'a'
					p = int(c-base) + 1
				} else if c >= 'A' && c <= 'Z' {
					var base rune = 'A'
					p = int(c-base) + 27
				}

			}
		}

	}
	total += p
	return total
}
