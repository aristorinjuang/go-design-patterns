package main

import "fmt"

type Command interface {
	Execute()
}

type ConcreteCommand struct {
	message string
}

func (c *ConcreteCommand) Execute() {
	fmt.Println(c.message)
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(c Command) {
	i.command = c
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	concreteCommand := &ConcreteCommand{message: "Hello World!"}
	invoker := &Invoker{}
	invoker.SetCommand(concreteCommand)
	invoker.ExecuteCommand()
}
