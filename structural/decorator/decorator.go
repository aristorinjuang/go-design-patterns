package main

import "fmt"

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (cc *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

type Decorator interface {
	Component
	SetComponent(Component)
}

type ConcreteDecoratorA struct {
	component Component
}

func (cda *ConcreteDecoratorA) SetComponent(component Component) {
	cda.component = component
}

func (cda *ConcreteDecoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", cda.component.Operation())
}

type ConcreteDecoratorB struct {
	component Component
}

func (cdb *ConcreteDecoratorB) SetComponent(component Component) {
	cdb.component = component
}

func (cdb *ConcreteDecoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)", cdb.component.Operation())
}

func main() {
	component := &ConcreteComponent{}
	decoratorA := &ConcreteDecoratorA{}
	decoratorB := &ConcreteDecoratorB{}

	decoratorA.SetComponent(component)
	decoratorB.SetComponent(decoratorA)

	result := decoratorB.Operation()
	fmt.Println(result)
}
