package main

import "fmt"

func main() {
	var length, capacity int

	fmt.Scanln(&length, &capacity)
	numbers := make([]int, length, capacity)

	fmt.Printf("len=%d cap=%d elements=%v", length, capacity, numbers)
}
