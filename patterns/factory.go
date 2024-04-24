package patterns

// Product представляет интерфейс создаваемых продуктов.
type Product interface {
	Use() string
}

// ConcreteProduct1 представляет конкретный продукт №1.
type ConcreteProduct1 struct{}

func (p *ConcreteProduct1) Use() string {
	return "Using ConcreteProduct1"
}

// ConcreteProduct2 представляет конкретный продукт №2.
type ConcreteProduct2 struct{}

func (p *ConcreteProduct2) Use() string {
	return "Using ConcreteProduct2"
}

// Creator представляет абстрактный класс с фабричным методом.
type Creator interface {
	FactoryMethod() Product
}

// ConcreteCreator1 представляет конкретный создатель №1.
type ConcreteCreator1 struct{}

func (c *ConcreteCreator1) FactoryMethod() Product {
	return &ConcreteProduct1{}
}

// ConcreteCreator2 представляет конкретный создатель №2.
type ConcreteCreator2 struct{}

func (c *ConcreteCreator2) FactoryMethod() Product {
	return &ConcreteProduct2{}
}
