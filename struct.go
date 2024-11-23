package main

import "fmt"

func main() {
	//1.
	type person struct {
		name      string
		age       int
		favcolors []string
	}
	me := person{name: "hrutja", age: 21, favcolors: []string{"black", "white"}}
	fmt.Printf("my info %v\n", me)
	you := person{name: "john", age: 25, favcolors: []string{"brown", "blue"}}
	fmt.Printf("your info %+v\n", you)

	//2.
	me.name = "andrei"
	var favcolors []string = you.favcolors
	fmt.Printf("type: %T, value: %v\n", favcolors, favcolors)
	favcolors = append(favcolors, "green")
	you.favcolors = favcolors
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

	//3.
	for index, favcolors := range me.favcolors {
		fmt.Printf("%d -> %q\n", index, favcolors)
	}

	//4

	type person struct {
		name   string
		age    int
		colors []string
		gr     grades
	}
	
	// 1.
	type grades struct {
		grade  int
		course string
	}

		// 3.
		me := person{
			name:   "Marius",
			age:    30,
			colors: []string{"red", "yellow"},
			gr:     grades{grade: 85, course: "Python"},
		}
		you := person{
			name:   "Maria",
			age:    22,
			colors: []string{"pink", "blue"},
			gr:     grades{grade: 100, course: "History"},
		}
	
		// 4.
		me.gr.grade = 98
		me.gr.course = "Golang"
	
		// 5.
		fmt.Printf("%v\n", me)
		fmt.Printf("%+v\n", you)
	
	}
}
	