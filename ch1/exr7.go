package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	result := 0

	for num > 0 {
		digit := num % 10
		num /= 10
		result = result*10 + digit
	}

	fmt.Println(result)
}
