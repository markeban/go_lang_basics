package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	joe := Person{name: "Joe", age: 35}
	fmt.Println(joe.name)
	fmt.Println(joe.age)
}
