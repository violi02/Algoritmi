/* IMPLEMENTAZIONE DEI 3 ALGORITMI DI ORDINAMENTO
INSERTION SORT, SELECTION SORT E MERGE SORT
*/

package main

import "fmt"

func main() {
	/*var num int
	n := 5
	vettore := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		fmt.Println(insertionSort(num, &vettore, len(vettore)))
	}*/
	prova := []int{15, 96, 44, 22, 54, 28, 83}
	fmt.Println(selectionSort(&prova, len(prova)))
}

func insertionSort(x int, v *[]int, l int) []int {
	if l == 0 {
		*v = append(*v, x)
		return *v
	}
	if (*v)[l-1] > x {
		spostare := (*v)[l-1]
		(*v)[l-1] = x
		*v = append(*v, spostare)
	} else {
		*v = append(*v, x)
	}
	return *v
}

func selectionSort(v *[]int, l int) *[]int {
	var k int
	inizio := (*v)[0]

	for i := 1; i < l-1; i++ {
		fmt.Println((*v)[i])
		fmt.Println("fine ", (*v)[l-i])
		if inizio > (*v)[l-i] {
			(*v)[k] = (*v)[l-i]
			fmt.Println((*v)[k], (*v)[l-i])
			(*v)[l-i] = inizio
		}
		k = k + 1
		inizio = (*v)[k]

	}
	return v
}

func mergeSort(v []int) []int {
	if len(v) == 0 || len(v) == 1 {
		return v
	}
	m := (len(v)) / 2
	v1 := v[:m]
	v2 := v[m:]

	v1 = mergeSort(v1)
	v2 = mergeSort(v2)

	return merge(v1, v2)
}

func merge(v1, v2 []int) []int {
	i1, i2 := 0, 0
	new := make([]int, 0)
	for i1 < len(v1) && i2 < len(v2) {
		if v1[i1] < v2[i1] {
			new = append(new, v1[i1])
			i1++
		}
		if v1[i1] > v2[i1] {
			new = append(new, v2[i2])
			i2++
		}
	}
	if i1 == len(v1) {
		for i := i2; i < len(v2); i++ {
			new = append(new, v2[i])
		}
	} else if i2 == len(v2) {
		for i := i1; i < len(v1); i++ {
			new = append(new, v1[i])
		}
	}
	return new
}
