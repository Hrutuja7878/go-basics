//Coding Exercise #2
//Change the code from the previous exercise and use the continue statement to print out all the numbers divisible by 7 between 1 and 50.

package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%d", i)
	}

	fmt.Println("")
}
