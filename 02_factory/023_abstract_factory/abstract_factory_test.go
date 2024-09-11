package main

import (
	"testing"
)

// 抽象工厂模式的单元测试
func TestAdidasFactory_MakeShoe(t *testing.T) {
	// 测试 Adidas 工厂创建的鞋子
	adidasFactory := &AdidasFactory{}
	adidasShoe := adidasFactory.MakeShoe()
	if adidasShoe.GetShoeType() != "Adidas Shoe" {
		t.Errorf("Expected Adidas Shoe, got %s", adidasShoe.GetShoeType())
	}
}

func TestAdidasFactory_MakeShirt(t *testing.T) {
	// 测试 Adidas 工厂创建的衣服
	adidasFactory := &AdidasFactory{}
	adidasShirt := adidasFactory.MakeShirt()
	if adidasShirt.GetShirtType() != "Adidas Shirt" {
		t.Errorf("Expected Adidas Shirt, got %s", adidasShirt.GetShirtType())
	}
}

func TestNikeFactory_MakeShoe(t *testing.T) {
	// 测试 Nike 工厂创建的鞋子
	nikeFactory := &NikeFactory{}
	nikeShoe := nikeFactory.MakeShoe()
	if nikeShoe.GetShoeType() != "Nike Shoe" {
		t.Errorf("Expected Nike Shoe, got %s", nikeShoe.GetShoeType())
	}
}

func TestNikeFactory_MakeShirt(t *testing.T) {
	// 测试 Nike 工厂创建的衣服
	nikeFactory := &NikeFactory{}
	nikeShirt := nikeFactory.MakeShirt()
	if nikeShirt.GetShirtType() != "Nike Shirt" {
		t.Errorf("Expected Nike Shirt, got %s", nikeShirt.GetShirtType())
	}
}
