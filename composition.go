package main

import "fmt"

type Animal struct {
	Diet  string
	Genus string
}

type Person struct {
	name string
	age  int
	Animal
	Genus string
	Diet  bool
}

func (a Animal) eat() bool {
	fmt.Println("im eating...")
	return true
}

func main() {
	joe := Person{
		name: "Joe",
		age:  35,
		Diet: false,
		Animal: Animal{
			Diet:  "ominvore",
			Genus: "primate",
		},
	}
	fmt.Println(joe.name)
	fmt.Println(joe.age)
	fmt.Println(joe.Diet)
	fmt.Println(joe.Genus)
	joe.eat()
}
