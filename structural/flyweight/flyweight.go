package main

import "fmt"

type Flyweight interface {
	Operation()
}

type ConcreteFlyweight struct {
	state string
}

func (f *ConcreteFlyweight) Operation() {
	fmt.Printf("ConcreteFlyweight: %s\n", f.state)
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := f.flyweights[key]; ok {
		return flyweight
	}
	flyweight := &ConcreteFlyweight{state: key}
	f.flyweights[key] = flyweight
	return flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

func main() {
	factory := NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("key1")
	flyweight2 := factory.GetFlyweight("key2")
	flyweight3 := factory.GetFlyweight("key1")

	flyweight1.Operation()
	flyweight2.Operation()
	flyweight3.Operation()

	fmt.Println(len(factory.flyweights))
}
