package builder

type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetResult() Product
}

type Product struct {
	parts []string
}

func (p *Product) AddPart(part string) {
	p.parts = append(p.parts, part)
}

type ConcreteBuilder1 struct {
	product Product
}

func (b *ConcreteBuilder1) BuildPart1() {
	b.product.AddPart("Part 1 of ConcreteBuilder1")
}

func (b *ConcreteBuilder1) BuildPart2() {
	b.product.AddPart("Part 2 of ConcreteBuilder1")
}

func (b *ConcreteBuilder1) BuildPart3() {
	b.product.AddPart("Part 3 of ConcreteBuilder1")
}

func (b *ConcreteBuilder1) GetResult() Product {
	return b.product
}

type ConcreteBuilder2 struct {
	product Product
}

func (b *ConcreteBuilder2) BuildPart1() {
	b.product.AddPart("Part 1 of ConcreteBuilder2")
}

func (b *ConcreteBuilder2) BuildPart2() {
	b.product.AddPart("Part 2 of ConcreteBuilder2")
}

func (b *ConcreteBuilder2) BuildPart3() {
	b.product.AddPart("Part 3 of ConcreteBuilder2")
}

func (b *ConcreteBuilder2) GetResult() Product {
	return b.product
}

type Director struct{}

func (d Director) Construct(builder Builder) Product {
	builder.BuildPart1()
	builder.BuildPart2()
	builder.BuildPart3()
	return builder.GetResult()
}
