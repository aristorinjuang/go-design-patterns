package main

import "fmt"

type Component interface {
	Operation()
}

type Leaf struct{}

func (leaf *Leaf) Operation() {
	fmt.Println("Leaf Operation")
}

type Composite struct {
	children []Component
}

func (composite *Composite) Add(child Component) {
	composite.children = append(composite.children, child)
}

func (composite *Composite) Remove(child Component) {
	for i, c := range composite.children {
		if c == child {
			composite.children = append(composite.children[:i], composite.children[i+1:]...)
			return
		}
	}
}

func (composite *Composite) Operation() {
	fmt.Println("Composite Operation")
	for _, child := range composite.children {
		child.Operation()
	}
}

func main() {
	root := &Composite{}
	root.Add(&Leaf{})
	branch := &Composite{}
	branch.Add(&Leaf{})
	branch.Add(&Leaf{})
	root.Add(branch)
	root.Operation()
}
