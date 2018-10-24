package main

import "fmt"

func main() {
	answer := is_cool("Frank")
	fmt.Println(answer)
}

func is_cool(name string) bool {
	if name == "Frank" {
		return true
	}
	return false
}
