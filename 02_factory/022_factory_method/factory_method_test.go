package main

import (
	"testing"
)

// 工厂方法模式的单元测试
func TestConcreteFactoryA_CreateProduct(t *testing.T) {
	// 测试工厂 A 创建的产品 A
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	if productA.GetName() != "Product A" {
		t.Errorf("Expected Product A, got %s", productA.GetName())
	}
}

func TestConcreteFactoryB_CreateProduct(t *testing.T) {
	// 测试工厂 B 创建的产品 B
	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	if productB.GetName() != "Product B" {
		t.Errorf("Expected Product B, got %s", productB.GetName())
	}
}
