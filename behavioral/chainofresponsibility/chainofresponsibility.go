package main

import "fmt"

type Handler interface {
	HandleRequest(request int)
	SetNextHandler(handler Handler)
}

type ConcreteHandler1 struct {
	nextHandler Handler
}

func (c *ConcreteHandler1) HandleRequest(request int) {
	switch {
	case request >= 0 && request < 10:
		fmt.Printf("Request %d handled by ConcreteHandler1\n", request)
	case c.nextHandler != nil:
		c.nextHandler.HandleRequest(request)
	default:
		fmt.Println("Request cannot be handled by any handler")
	}
}

func (c *ConcreteHandler1) SetNextHandler(handler Handler) {
	c.nextHandler = handler
}

type ConcreteHandler2 struct {
	nextHandler Handler
}

func (c *ConcreteHandler2) HandleRequest(request int) {
	switch {
	case request >= 10 && request < 20:
		fmt.Printf("Request %d handled by ConcreteHandler2\n", request)
	case c.nextHandler != nil:
		c.nextHandler.HandleRequest(request)
	default:
		fmt.Println("Request cannot be handled by any handler")
	}
}

func (c *ConcreteHandler2) SetNextHandler(handler Handler) {
	c.nextHandler = handler
}

type Client struct {
	handler Handler
}

func (c *Client) MakeRequest(request int) {
	c.handler.HandleRequest(request)
}

func main() {
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}

	handler1.SetNextHandler(handler2)

	client := &Client{handler: handler1}

	client.MakeRequest(5)
	client.MakeRequest(15)
	client.MakeRequest(25)
}
