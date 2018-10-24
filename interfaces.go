package main

type Jedi interface {
	HasForce() bool
}

type Sith struct {
}

type Knight struct {
}

func (s Sith) HasForce() bool {
	name = "Bob"
	return true
}

func (k Knight) HasForce() bool {
	return true
}

func main() {
	var anakin Jedi = Knight{}
	anakin.HasForce()
}
