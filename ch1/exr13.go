package main

import (
	"errors"
	"fmt"
	"math"
)

func negativeNumError() error {
	err := errors.New("math: can't calculate square root of negative number")
	return err
}

func main() {
	var num float64
	fmt.Scanf("%f", &num)

	if num < 0 {
		err := negativeNumError()
		fmt.Println(err)
		return
	}
	fmt.Println(math.Sqrt(num))
}
