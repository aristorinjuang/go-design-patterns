package main

import (
	"fmt"
)

type Expression interface {
	Interpret() int
}

type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

type Add struct {
	left  Expression
	right Expression
}

func (a *Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

type Subtract struct {
	left  Expression
	right Expression
}

func (s *Subtract) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

type Multiply struct {
	left  Expression
	right Expression
}

func (m *Multiply) Interpret() int {
	return m.left.Interpret() * m.right.Interpret()
}

type Divide struct {
	left  Expression
	right Expression
}

func (d *Divide) Interpret() int {
	return d.left.Interpret() / d.right.Interpret()
}

func main() {
	expression := &Add{
		left: &Number{value: 10},
		right: &Multiply{
			left:  &Number{value: 2},
			right: &Subtract{left: &Number{value: 5}, right: &Number{value: 3}},
		},
	}
	result := expression.Interpret()
	fmt.Printf("Result: %d\n", result)
}
