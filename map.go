package main

import "fmt"

func main() {
	//1.
	var m1 map[string]int
	fmt.Printf("m1 type: %T, m1 value : %#v\n ", m1, m1)

	m2 := map[string]float64{"hrutuja": 11.2}

	m2["abba"] = 11.3
	fmt.Println(m2)
	fmt.Println(m2["hrutuja"])
	fmt.Println(m2["john"])

	//3.
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 1)
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("k: %d, v: %t\n", k, v)
	}

}
