package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype1 struct {
	property1 string
	property2 int
}

func (p *ConcretePrototype1) Clone() Prototype {
	return &ConcretePrototype1{p.property1, p.property2}
}

type ConcretePrototype2 struct {
	property1 int
	property2 int
}

func (p *ConcretePrototype2) Clone() Prototype {
	return &ConcretePrototype2{p.property1, p.property2}
}

type Client struct{}

func (c *Client) Operation(prototype Prototype) {
	clone := prototype.Clone()
	fmt.Printf("Cloned object: %+v\n", clone)
}

func main() {
	prototype1 := &ConcretePrototype1{"hello", 123}
	client := &Client{}
	client.Operation(prototype1)

	prototype2 := &ConcretePrototype2{456, 789}
	client.Operation(prototype2)
}
