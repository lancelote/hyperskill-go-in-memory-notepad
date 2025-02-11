package main

import "fmt"

func main() {
	var v int

	_, err := fmt.Scan(&v)
	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}
