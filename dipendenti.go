package main

import "fmt"

func main() {

	dipendenti := make(map[string]string)
	// mappa dipendenti associazione subordinato -- supervisore
	dipendenti["Bruno"] = "Anna"
	dipendenti["Carlo"] = "Anna"
	dipendenti["Daniela"] = "Anna"
	dipendenti["Enrico"] = "Bruno"
	dipendenti["Franci"] = "Bruno"
	dipendenti["Irene"] = "Franci"
	dipendenti["Harry"] = "Gianni"

	fmt.Println("Subordinati di Anna? ")
	stampaSubordinati("Anna", dipendenti)
	fmt.Println("Subordinati di Bruno? ")
	stampaSubordinati("Bruno", dipendenti)
	fmt.Println("Subordinati di Franci? ")
	stampaSubordinati("Franci", dipendenti)
	fmt.Println("Subordinati di Gianni? ")
	stampaSubordinati("Gianni", dipendenti)

	for k, _ := range dipendenti {
		supervisore(k, dipendenti)

	}
	for k, _ := range dipendenti {
		fmt.Print("Gli impiegati sopra ", k, " sono ")
		stampaImpiegatiSopra(k, dipendenti)
		fmt.Println()
	}
	fmt.Println("Impiegati con supervisore : ")
	for k, _ := range dipendenti {
		stampaImpiegatiConSupervisore(k, dipendenti)
	}
	fmt.Println("Impiegati senza subordinati:")
	for k, _ := range dipendenti {
		quantiSenzaSubordinati(k, dipendenti)
	}
}

func stampaSubordinati(supervisore string, dipendenti map[string]string) {
	for k, val := range dipendenti {
		if val == supervisore {
			fmt.Println(k)
		}
	}
}

// dato un sub stampa supervisore
func supervisore(subordinato string, dipendenti map[string]string) {
	for k, val := range dipendenti {
		if k == subordinato {
			fmt.Println("Il supervisore di", subordinato, "è", val)
		}
	}
}

// DA RIGUARDARE
func stampaImpiegatiSopra(impiegato string, dipendenti map[string]string) {
	// associare ad ogni dipende un livello
	// fin quando impiegato è uguale ad "Anna"
	for {
		if dipendenti[impiegato] == "" {
			fmt.Println(impiegato, "non ha dipendenti superiori")
			return
		}
		fmt.Print(dipendenti[impiegato])
		padre := dipendenti[impiegato]
		impiegato = padre
	}

}

// DA FARE
func quantiSenzaSubordinati(impiegato string, dipendenti map[string]string) {
	// se il mio impiegato non compare nella lista degli elementi
	for k, v := range dipendenti {

	}

}
func stampaImpiegatiConSupervisore(impiegato string, dipendenti map[string]string) {
	for k, _ := range dipendenti {
		if k == impiegato {
			if _, ok := dipendenti[k]; ok {
				fmt.Println(impiegato)
			}
		}
	}
}
