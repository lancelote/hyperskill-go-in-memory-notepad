package main

import "fmt"

func solve(ptr *int) int {
	var num = *ptr
	return num
}

func main() {
	var num = new(int)
	fmt.Scan(num)

	var test = solve(num)
	fmt.Println(test)
}
