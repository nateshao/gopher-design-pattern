package main

import "fmt"

// Product 是产品接口
type Product interface {
	GetName() string
}

// ConcreteProductA 是具体产品 A
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "Product A"
}

// ConcreteProductB 是具体产品 B
type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "Product B"
}

// Factory 是工厂接口，定义了创建产品的方法
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA 是具体工厂 A
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB 是具体工厂 B
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// 使用具体工厂 A 创建产品 A
	var factory Factory = &ConcreteFactoryA{}
	productA := factory.CreateProduct()
	fmt.Println(productA.GetName()) // 输出：Product A

	// 使用具体工厂 B 创建产品 B
	factory = &ConcreteFactoryB{}
	productB := factory.CreateProduct()
	fmt.Println(productB.GetName()) // 输出：Product B
}
