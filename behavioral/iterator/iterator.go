package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	index     int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.aggregate.items)
}

func (i *ConcreteIterator) Next() interface{} {
	if !i.HasNext() {
		return nil
	}
	item := i.aggregate.items[i.index]
	i.index++
	return item
}

type ConcreteAggregate struct {
	items []interface{}
}

func (a *ConcreteAggregate) CreateIterator() Iterator {
	return &ConcreteIterator{
		aggregate: a,
		index:     0,
	}
}

func main() {
	aggregate := &ConcreteAggregate{
		items: []interface{}{"A", "B", "C", "D", "E"},
	}
	iterator := aggregate.CreateIterator()
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Printf("%s ", item)
	}
}
