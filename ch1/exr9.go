package main

import "fmt"

func getSumOfVectorElements(vector []int) int {
	var sum = 0

	for _, x := range vector {
		sum += x
	}

	return sum
}

func getProductOfVectorElements(vector []int) int {
	var product = 1

	for _, x := range vector {
		product *= x
	}

	return product
}

func main() {
	var vc []int = []int{1, 2, 3, 4, 5, 6}

	var sum, prod int
	
	sum = getSumOfVectorElements(vc)
	prod = getProductOfVectorElements(vc)

	fmt.Println(sum, prod)
}
