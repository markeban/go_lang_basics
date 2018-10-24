package main

import (
	"fmt"
)

type animate interface {
	walk()
}

type monkey struct {
	name string
}

type human struct {
	name string
	age  int
}

func (h human) walk() {
	fmt.Println(h)
	fmt.Println("The human is walking.")
}

func (h human) addLastName(lastName string) string {
	fullName := fmt.Sprintf("%s %s", h.name, lastName)
	return fullName
}

func main() {
	bob := human{name: "Bob", age: 50}
	bob.walk()
	fmt.Println(bob.addLastName("McCready"))
}
