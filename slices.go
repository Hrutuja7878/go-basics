package main

import "fmt"

func main() {
	//6
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

	//7
	people := []string{"Marry", "John", "Paul", "Diana"}
	newpeople := []string{}
	newpeople = append(newpeople, people...)
	newpeople[1] = "Dan"
	fmt.Println(people, newpeople)

	//8
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Println(newYears)

}
