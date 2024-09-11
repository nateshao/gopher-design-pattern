package main

import "testing"

// 简单工厂模式的单元测试
func TestCreateProduct(t *testing.T) {
	// 测试创建产品 A
	productA := CreateProduct("A")
	if productA.GetName() != "concrete Product A" {
		t.Errorf("Expected Product A, got %s", productA.GetName())
	}

	// 测试创建产品 B
	productB := CreateProduct("B")
	if productB.GetName() != "concrete Product B" {
		t.Errorf("Expected Product B, got %s", productB.GetName())
	}

	// 测试无效产品类型
	productInvalid := CreateProduct("C")
	if productInvalid != nil {
		t.Errorf("Expected nil, got %v", productInvalid)
	}
}
