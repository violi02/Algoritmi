package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SOLO PARTE 1
func main() {
	// lettura da file

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		os.Exit(-1)
	}
	defer file.Close()

	// creo mappa associazione padre - figlio, elementi non in ordine
	orbita := make(map[string]string)
	// prendo dal file di input associazione padre - figlio
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		valori := strings.Split(str, ")")
		orbita[valori[1]] = valori[0]
	}

	var numeroIndirette, numeroDirette int
	// calcolo di orbite indirette
	for key, _ := range orbita {
		numeroIndirette += orbiteInd(key, &orbita)
		numeroDirette++
	}
	fmt.Println("Numero di orbite indirette:", numeroIndirette)
	fmt.Println("Numero di orbite dirette:", numeroDirette)
	fmt.Println("Numero di orbite totali:", numeroIndirette+numeroDirette)

}

// per avere l'altezza dell'albero devo prendere tutti gli elementi che non si ripetono
func orbiteInd(figlio string, orbita *map[string]string) int {
	padre := (*orbita)[figlio]
	var counter int

	for padre != "COM" {
		counter++
		nonno := (*orbita)[padre]
		figlio = padre
		padre = nonno
	}
	return counter
}
