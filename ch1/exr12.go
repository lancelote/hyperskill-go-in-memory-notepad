package main

import "fmt"

func main() {
	elements := []string{"Helium", "Boron", "Neon", "Calcium", "Iodine"}

	elem1 := elements[0][:2]
	elem2 := elements[1][:1]
	elem3 := elements[2][:2]
	elem4 := elements[3][:2]
	elem5 := elements[4][:1]

	fmt.Println(elem1, elem2, elem3, elem4, elem5)
}
