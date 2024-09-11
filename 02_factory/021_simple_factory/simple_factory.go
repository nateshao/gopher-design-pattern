package main

import "fmt"

// Product 是一个接口，所有产品类型都实现这个接口
type Product interface {
	GetName() string
}

// ConcreteProductA 是具体产品 A
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "concrete Product A"
}

// ConcreteProductB 是具体产品 B
type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "concrete Product B"
}

// 工厂函数，根据输入的类型返回对应的产品实例
func CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	// 使用工厂创建产品 A
	productA := CreateProduct("A")
	if productA != nil {
		fmt.Println(productA.GetName()) // concrete Product A
	}

	// 使用工厂创建产品 B
	productB := CreateProduct("B")
	if productB != nil {
		fmt.Println(productB.GetName()) // 输出：concrete Product B
	}
}
