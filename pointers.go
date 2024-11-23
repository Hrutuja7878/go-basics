package main

import "fmt"

func swap(p, q *float64) {
	*p, *q = *q, *p
}

func main() {
	//1.
	x := 10.10
	fmt.Println("address of x is \n", &x)
	ptr := &x
	fmt.Println(ptr)
	fmt.Printf("type is: %T, value is: %v\n", ptr, ptr)
	fmt.Printf("the address of of ptr: %p\n", &ptr)
	fmt.Printf("the value of ptr is: %f\n", *ptr)

	//2.
	a, b := 10, 2
	ptra, ptrb := &a, &b
	c := *ptra / *ptrb
	fmt.Printf("c is %d\n", c)

	//3.
	y, z := 5.5, 8.8
	swap(&y, &z)
	fmt.Printf("y is %v and z ia %v\n", y, z)

}
