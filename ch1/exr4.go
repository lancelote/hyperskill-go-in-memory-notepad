package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)

	if number > 0 {
		fmt.Println("Positive!")
	} else if number < 0 {
		fmt.Println("Negative!")
	} else {
		fmt.Println("Zero!")
	}
}
