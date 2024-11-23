package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [4]int{1, 2, 3, 4}
	fmt.Printf("%#v\n", numbers)

	numbers[1] = 8
	fmt.Printf("%#v\n", numbers)

	// numbers[5] = 10 // error

	for i, v := range numbers {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value:", numbers[i])
	}
	// multidimensional array
	balance := [2][3]int{
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(balance)

	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m)
	m[0] = 2
	fmt.Println("n is equal to m:", n == m)
}
