//Coding Exercise #4

//Using a for loop, an if statement and the logical and operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.

package main

import "fmt"

func main() {
for i := 2003; i <= 2024; i++ {
	fmt.Printf("%d ", i)

}
fmt.Println()
}

package main

import "fmt"

func main() {
	birthYear, currentYear := 2003, 2024

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Printl()
}
