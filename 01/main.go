package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) GetName() string {
	return h.name
}

type Action struct {
	Human
}

func main() {
	h := Human{name: "Bob"}

	a := Action{Human: h}

	fmt.Println(a.GetName())

}
