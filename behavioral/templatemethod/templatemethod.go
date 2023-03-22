package main

import "fmt"

type AbstractClass interface {
	TemplateMethod()
	PrimitiveOperation1()
	PrimitiveOperation2()
}

type ConcreteClass struct{}

func (c *ConcreteClass) TemplateMethod() {
	fmt.Println("ConcreteClass: template method starts")
	c.PrimitiveOperation1()
	c.PrimitiveOperation2()
	fmt.Println("ConcreteClass: template method ends")
}

func (c *ConcreteClass) PrimitiveOperation1() {
	fmt.Println("ConcreteClass: primitive operation 1")
}

func (c *ConcreteClass) PrimitiveOperation2() {
	fmt.Println("ConcreteClass: primitive operation 2")
}

func main() {
	concrete := &ConcreteClass{}
	concrete.TemplateMethod()
}
