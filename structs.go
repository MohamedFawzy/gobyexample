package main

import "fmt"

type Person struct {

	name string
	age int
}

func main() {

	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{name: "Alex", age:30})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{name: "Ann", age:40})

	s := Person{name: "Sean", age:37}

	fmt.Println(s)

	sp := &s

	fmt.Println(sp.age)

	sp.age=44

	fmt.Println(sp.age)
}
