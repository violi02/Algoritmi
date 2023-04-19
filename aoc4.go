package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	leggiI()
}

func leggiI() {
	file, err := os.Open("quarto.txt")
	if err != nil {
		fmt.Println("errore nella lettura del file")
		os.Exit(-1)
	}

	var totale_fully, totale_part int
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// lettura input es. 2-4,6-8...
		input := strings.Split(scanner.Text(), ",")
		s1 := strings.Split(input[0], "-")
		s2 := strings.Split(input[1], "-")
		fully_contain(s1, s2)
		if fully_contain(s1, s2) {
			totale_fully = totale_fully + 1
		}
		if overlap(s1, s2) {
			totale_part = totale_part + 1
		}

	}
	fmt.Println(totale_fully)
	fmt.Println(totale_part)
}

//ogni riga

func fully_contain(s1, s2 []string) bool {
	//es. s1 = 2-4, s2 = 6-8
	inizio_1, _ := strconv.Atoi(s1[0])
	fine_1, _ := strconv.Atoi(s1[1])
	inizio_2, _ := strconv.Atoi(s2[0])
	fine_2, _ := strconv.Atoi(s2[1])

	// controllo se primo range è contenuto completemante in secondo
	if inizio_1 < inizio_2 {
		for i := inizio_1; i <= fine_1; i++ {
			if i == fine_2 {
				// vuol dire che il mio secondo range è contenuto
				// ritorno true
				return true
				break
			}
		}
	}
	// controllo se secondo range è contenuto completemante in secondo
	if inizio_1 > inizio_2 {
		for i := inizio_2; i <= fine_2; i++ {
			if i == fine_1 {
				// vuol dire che il mio secondo range è contenuto
				// ritorno true
				return true
				break
			}
		}
	}

	if inizio_1 == inizio_2 && fine_1 != fine_2 {
		if fine_1 > fine_2 {
			for i := inizio_1; i <= fine_1; i++ {
				if i == fine_2 {
					// vuol dire che il mio secondo range è contenuto
					// ritorno true
					return true
					break
				}
			}
		}
		if fine_1 < fine_2 {
			for i := inizio_2; i <= fine_2; i++ {
				if i == fine_1 {
					// vuol dire che il mio secondo range è contenuto
					// ritorno true
					return true
					break
				}
			}
		}
	}
	if inizio_1 == inizio_2 && fine_1 == fine_2 {
		return true
	}

	return false
}

func overlap(s1, s2 []string) bool {
	inizio_1, _ := strconv.Atoi(s1[0])
	fine_1, _ := strconv.Atoi(s1[1])
	inizio_2, _ := strconv.Atoi(s2[0])
	fine_2, _ := strconv.Atoi(s2[1])

	if inizio_1 == inizio_2 {
		return true
	}
	if fine_1 == fine_2 {
		return true
	}

	if inizio_1 < inizio_2 {
		for i := inizio_1; i <= fine_1; i++ {
			if i == inizio_2 {
				return true
			}
		}
	}
	if inizio_1 > inizio_2 {
		for i := inizio_2; i <= fine_2; i++ {
			if i == inizio_1 {
				return true
			}
		}
	}

	return false
}
