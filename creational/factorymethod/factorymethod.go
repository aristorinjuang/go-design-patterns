package factorymethod

import "fmt"

type Creator interface {
	FactoryMethod() Product
}

type Product interface {
	Use()
}

type ConcreteProduct1 struct{}

func (p *ConcreteProduct1) Use() {
	fmt.Println("Using ConcreteProduct1")
}

type ConcreteProduct2 struct{}

func (p *ConcreteProduct2) Use() {
	fmt.Println("Using ConcreteProduct2")
}

type ConcreteCreator1 struct{}

func (c *ConcreteCreator1) FactoryMethod() Product {
	return &ConcreteProduct1{}
}

type ConcreteCreator2 struct{}

func (c *ConcreteCreator2) FactoryMethod() Product {
	return &ConcreteProduct2{}
}
