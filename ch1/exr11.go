package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)
	intSlice := []int{num1, num2, num3}

	for _, x := range intSlice {
		fmt.Println(x * 2)
	}
}
