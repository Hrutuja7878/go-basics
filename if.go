// flow control
// Coding Exercise #1
// Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.

package main

import "fmt"

func main() {
	for i := 0; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
	fmt.Println("")
}
