package main

import "fmt"

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Handle request with ConcreteStateA")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Handle request with ConcreteStateB")
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
}

func main() {
	context := &Context{}
	stateA := &ConcreteStateA{}
	stateB := &ConcreteStateB{}

	context.SetState(stateA)
	context.Request()

	context.SetState(stateB)
	context.Request()
}
