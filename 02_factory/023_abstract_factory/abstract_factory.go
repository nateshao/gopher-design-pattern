package main

import "fmt"

// 产品族中的产品接口
type Shoe interface {
	GetShoeType() string
}

type Shirt interface {
	GetShirtType() string
}

// Adidas 鞋子和衣服实现
type AdidasShoe struct{}

func (s *AdidasShoe) GetShoeType() string {
	return "Adidas Shoe"
}

type AdidasShirt struct{}

func (s *AdidasShirt) GetShirtType() string {
	return "Adidas Shirt"
}

// Nike 鞋子和衣服实现
type NikeShoe struct{}

func (s *NikeShoe) GetShoeType() string {
	return "Nike Shoe"
}

type NikeShirt struct{}

func (s *NikeShirt) GetShirtType() string {
	return "Nike Shirt"
}

// 抽象工厂接口，定义创建产品族的方法
type SportsFactory interface {
	MakeShoe() Shoe
	MakeShirt() Shirt
}

// 具体工厂 AdidasFactory，负责创建 Adidas 产品族
type AdidasFactory struct{}

func (f *AdidasFactory) MakeShoe() Shoe {
	return &AdidasShoe{}
}

func (f *AdidasFactory) MakeShirt() Shirt {
	return &AdidasShirt{}
}

// 具体工厂 NikeFactory，负责创建 Nike 产品族
type NikeFactory struct{}

func (f *NikeFactory) MakeShoe() Shoe {
	return &NikeShoe{}
}

func (f *NikeFactory) MakeShirt() Shirt {
	return &NikeShirt{}
}

func main() {
	// Adidas 工厂
	adidasFactory := &AdidasFactory{}
	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()
	fmt.Println(adidasShoe.GetShoeType())   // 输出：Adidas Shoe
	fmt.Println(adidasShirt.GetShirtType()) // 输出：Adidas Shirt

	// Nike 工厂
	nikeFactory := &NikeFactory{}
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()
	fmt.Println(nikeShoe.GetShoeType())   // 输出：Nike Shoe
	fmt.Println(nikeShirt.GetShirtType()) // 输出：Nike Shirt
}
