package main

import "fmt"

//1.
func cube(n float64) float64 {
	return n * n * n
}

func main() {
	x := cube(3)
	fmt.Println(x)

	f, s := f1(4)
	fmt.Println("f:", f, "s:", s)

	f, s = f1(10)
	fmt.Println("f:", f, "s:", s)
}

//2.
func f1(n uint) (int, int) {
	fact := 1
	sum := 0
	for i := 1; i <= int(n); i++ {
		fact *= i
		sum += i
	}
	return fact, sum
}
