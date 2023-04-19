package main

import (
	"bufio"
	"fmt"
	"os"
)

func leggiInput() {
	file, err := os.Open("otto.txt")
	if err != nil {
		fmt.Println("Errore nell'apertura del file")
		os.Exit(-1)
	}
	defer file.Close()
	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		riga := scanner.Text()
		input = append(input, riga)
	}

	//righe, colonne := estraiDimensioni(input)
	var matrice [5][5]string
	// creazione della mia matrice
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			(matrice)[i][j] = string(input[i][j])
		}
	}
	myprint(matrice)
	alberivisibili := numeroAlberiVisibili(matrice)
	fmt.Println("alberivisibili:", alberivisibili)
}

func myprint(matrice [5][5]string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(matrice[i][j])
		}
		fmt.Println()
	}
}

// creo la mia matrice con l'altezza degli alberi
func estraiDimensioni(input []string) (int, int) {
	n_righe := len(input)
	n_colonne := len(input[0])
	return n_righe, n_colonne
}

func numeroAlberiVisibili(matrice [5][5]string) int {
	somma := 0
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			if matrice[i][j] > matrice[i-1][j] {
				somma += 1
				fmt.Println(i, j)
				fmt.Println("curr:", matrice[i][j], "riga prec:", matrice[i-1][j])
			} else if matrice[i][j] > matrice[i][j-1] {
				somma += 1
				fmt.Println(i, j)
				fmt.Println("curr:", matrice[i][j], "colonna prec:", matrice[i][j-1])
			} else if matrice[i][j] > matrice[i][j+1] {
				somma += 1
				fmt.Println(i, j)
				fmt.Println("curr:", matrice[i][j], "colonna dopo:", matrice[i][j+1])
			} else if matrice[i][j] > matrice[i+1][j] {
				somma += 1
				fmt.Println(i, j)
				fmt.Println("curr:", matrice[i][j], "riga dopo:", matrice[i+1][j])
			}
			break
		}
	}
	return somma
}

//------------MAIN-----------------
func main() {
	leggiInput()
}
