package main

import "fmt"

type Mediator interface {
	SendMessage(sender Colleague, message string)
}

type ConcreteMediator struct {
	colleague1 Colleague
	colleague2 Colleague
}

func (m *ConcreteMediator) SendMessage(sender Colleague, message string) {
	if sender == m.colleague1 {
		m.colleague2.Notify(message)
	} else {
		m.colleague1.Notify(message)
	}
}

type Colleague interface {
	SendMessage(message string)
	Notify(message string)
}

type ConcreteColleague1 struct {
	mediator Mediator
}

func (c *ConcreteColleague1) SendMessage(message string) {
	c.mediator.SendMessage(c, message)
}

func (c *ConcreteColleague1) Notify(message string) {
	fmt.Printf("ConcreteColleague1 received message: %s\n", message)
}

type ConcreteColleague2 struct {
	mediator Mediator
}

func (c *ConcreteColleague2) SendMessage(message string) {
	c.mediator.SendMessage(c, message)
}

func (c *ConcreteColleague2) Notify(message string) {
	fmt.Printf("ConcreteColleague2 received message: %s\n", message)
}

func main() {
	mediator := &ConcreteMediator{}

	colleague1 := &ConcreteColleague1{mediator: mediator}
	colleague2 := &ConcreteColleague2{mediator: mediator}

	mediator.colleague1 = colleague1
	mediator.colleague2 = colleague2

	colleague1.SendMessage("Hello, colleague2!")
	colleague2.SendMessage("Hi there, colleague1!")
}
