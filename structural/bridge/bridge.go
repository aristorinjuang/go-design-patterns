package main

import "fmt"

type Abstraction interface {
	Operation()
}

type ConcreteAbstractionA struct {
	implementation Implementation
}

func (aa *ConcreteAbstractionA) Operation() {
	aa.implementation.OperationImplementation()
}

type ConcreteAbstractionB struct {
	implementation Implementation
}

func (ab *ConcreteAbstractionB) Operation() {
	ab.implementation.OperationImplementation()
}

type Implementation interface {
	OperationImplementation()
}

type ConcreteImplementationA struct{}

func (cia *ConcreteImplementationA) OperationImplementation() {
	fmt.Println("ConcreteImplementationA operation")
}

type ConcreteImplementationB struct{}

func (cib *ConcreteImplementationB) OperationImplementation() {
	fmt.Println("ConcreteImplementationB operation")
}

func main() {
	concreteImplementationA := &ConcreteImplementationA{}
	concreteImplementationB := &ConcreteImplementationB{}

	concreteAbstractionA := &ConcreteAbstractionA{implementation: concreteImplementationA}
	concreteAbstractionB := &ConcreteAbstractionB{implementation: concreteImplementationB}

	concreteAbstractionA.Operation()
	concreteAbstractionB.Operation()
}
