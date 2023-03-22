package main

import "fmt"

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

func (e *ConcreteElementA) OperationA() {
	fmt.Println("ConcreteElementA: operation A")
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

func (e *ConcreteElementB) OperationB() {
	fmt.Println("ConcreteElementB: operation B")
}

type ConcreteVisitor1 struct{}

func (v *ConcreteVisitor1) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor1: visiting ConcreteElementA")
	element.OperationA()
}

func (v *ConcreteVisitor1) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor1: visiting ConcreteElementB")
	element.OperationB()
}

type ConcreteVisitor2 struct{}

func (v *ConcreteVisitor2) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor2: visiting ConcreteElementA")
	element.OperationA()
}

func (v *ConcreteVisitor2) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor2: visiting ConcreteElementB")
	element.OperationB()
}

func main() {
	elements := []Element{&ConcreteElementA{}, &ConcreteElementB{}}

	visitor1 := &ConcreteVisitor1{}
	visitor2 := &ConcreteVisitor2{}

	for _, element := range elements {
		element.Accept(visitor1)
		element.Accept(visitor2)
	}
}
