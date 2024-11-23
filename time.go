// operator and conversion
// Exercise 4
package main

import "fmt"

func main() {

	distance := 149600000 * 1000
	speed := 299792458

	time := distance / speed
	fmt.Printf("The sunlight reaches Earth in %v seaconds \n", time)

}
