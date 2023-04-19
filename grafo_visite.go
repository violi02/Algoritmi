// visite di un grafo, in base alla sua implementazione
package main

import "fmt"

type graph map[string][]string

func main() {

}

// visita in profondità-->ricorsiva
// per ottenerla iterativa utilizzo una pila di supporto
func dfs(g graph, v string, aux map[string]bool) {
	// arrivo al vertice e lo stampo
	fmt.Println(v)
	// vertice v visitato
	aux[v] = true
	// scorro gli adiacenti di v
	for _, v2 := range g[v] {
		// se non ho ancora visitato il nodo richiamo la funzione e lo visito
		if aux[v2] != true {
			dfs(g, v2, aux)
		}
	}
}

// per la visita in ampiezza utilizzo una coda
func bfs(g graph, v string, aux map[string]bool) {
	coda := []string{v}
	aux[v] = true

	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]
		fmt.Println(" ", v)

		for _, v2 := range g[v] {
			if aux[v2] != true {
				coda = append(coda, v2)
			}
		}
	}
}
