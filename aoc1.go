package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	num_calorie := letturaI()

	fmt.Println("numero massimo calorie:", num_calorie)
}

func letturaI() int {
	file, err := os.Open("primo.txt")
	if err != nil {
		fmt.Println("errore nella lettura del file")
		os.Exit(-1)
	}
	defer file.Close()
	i := 1
	var sum int
	// associa ad ogni elfo il suo quntitativo di calorie
	elfi_calorie := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		if input != "\t" {
			num_calorie, _ := strconv.Atoi(input)
			sum += num_calorie
			elfi_calorie[i] = sum
		}
		if input == "" {
			i++
			sum = 0
		}

	}

	ints := make([]int, 0)
	max := 0

	for _, val := range elfi_calorie {
		ints = append(ints, val)

	}
	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)
	for k := len(ints) - 1; k > len(ints)-4; k-- {
		max += ints[k]
	}
	return max
}
