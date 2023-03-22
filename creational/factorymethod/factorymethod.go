package factorymethod

type Creator interface {
	FactoryMethod() Product
}

type Product interface {
	Use()
}

type ConcreteProduct1 struct{}

func (p *ConcreteProduct1) Use() {}

type ConcreteProduct2 struct{}

func (p *ConcreteProduct2) Use() {}

type ConcreteCreator1 struct{}

func (c *ConcreteCreator1) FactoryMethod() Product {
	return &ConcreteProduct1{}
}

type ConcreteCreator2 struct{}

func (c *ConcreteCreator2) FactoryMethod() Product {
	return &ConcreteProduct2{}
}
