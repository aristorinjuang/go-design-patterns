package abstractfactory

import "fmt"

type ProductA interface {
	DoSomething()
}

type ProductB interface {
	DoSomethingElse()
}

type ConcreteProductA1 struct{}

func (p ConcreteProductA1) DoSomething() {
	fmt.Println("Doing something in ConcreteProductA1")
}

type ConcreteProductB1 struct{}

func (p ConcreteProductB1) DoSomethingElse() {
	fmt.Println("Doing something else in ConcreteProductB1")
}

type ConcreteProductA2 struct{}

func (p ConcreteProductA2) DoSomething() {
	fmt.Println("Doing something in ConcreteProductA2")
}

type ConcreteProductB2 struct{}

func (p ConcreteProductB2) DoSomethingElse() {
	fmt.Println("Doing something else in ConcreteProductB2")
}

type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ConcreteFactory1 struct{}

func (c ConcreteFactory1) CreateProductA() ProductA {
	return ConcreteProductA1{}
}

func (c ConcreteFactory1) CreateProductB() ProductB {
	return ConcreteProductB1{}
}

type ConcreteFactory2 struct{}

func (c ConcreteFactory2) CreateProductA() ProductA {
	return ConcreteProductA2{}
}

func (c ConcreteFactory2) CreateProductB() ProductB {
	return ConcreteProductB2{}
}

func GetFactory(factoryType string) AbstractFactory {
	switch factoryType {
	case "Factory1":
		return ConcreteFactory1{}
	case "Factory2":
		return ConcreteFactory2{}
	default:
		return nil
	}
}
