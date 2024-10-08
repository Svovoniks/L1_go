package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetAge() int {
	return h.age
}

type Action struct {
	Human
}

func main() {
	h := Human{
		name: "Bob",
		age:  100,
	}

	a := Action{Human: h}

	fmt.Println(a.GetName())
	fmt.Println(a.GetAge())
}
