package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	total := letturaInput()
	fmt.Println(total)
}

func letturaInput() int {
	file, err := os.Open("secondo.txt")
	if err != nil {
		fmt.Println("errore nella lettura del file")
		os.Exit(-1)
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		//sasso
		if input[0] == "A" {
			// pareggio scegliendo sasso
			if input[1] == "Y" {
				total += 4
			}
			// devo perdere --> forbici
			if input[1] == "X" {
				total += 3
			}
			//vittoria scegliendo paper
			if input[1] == "Z" {
				total += 8
			}
		}
		//carta
		if input[0] == "B" {

			// perdo se scelgo sasso
			if input[1] == "X" {
				total += 1
			}
			// vinco se scelgo forbici
			if input[1] == "Z" {
				total += 9
			}
			// pareggio con carta
			if input[1] == "Y" {
				total += 5
			}
		}
		//forbice
		if input[0] == "C" {
			// vinco se scelgo sass7
			if input[1] == "Z" {
				total += 7
			}
			// perdo se scelgo carta
			if input[1] == "X" {
				total += 2
			}
			// pareggio se scelgo forbici
			if input[1] == "Y" {
				total += 6
			}
		}
	}
	return total
}
