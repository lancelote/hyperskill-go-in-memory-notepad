package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	degree := 0
	value := 1

	for value < x {
		value *= 2
		degree++
	}

	fmt.Println(degree)
}
