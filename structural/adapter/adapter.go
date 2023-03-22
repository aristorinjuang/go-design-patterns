package main

import "fmt"

type Target interface {
	Request()
}

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() {
	fmt.Println("Adaptee's specific request")
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() {
	fmt.Println("Adapter's request")
	a.adaptee.SpecificRequest()
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func main() {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)

	adapter.Request()
}
