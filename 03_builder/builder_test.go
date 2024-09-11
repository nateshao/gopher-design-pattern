package main

import (
	"testing"
)

// 测试计算机建造
func TestComputerBuilder(t *testing.T) {
	builder := NewComputerBuilder()
	computer := builder.SetCPU("AMD Ryzen 9").
		SetGPU("NVIDIA RTX 3090").
		SetRAM(64).
		SetStorage(2048).
		SetDisplay("5K").
		Build()

	if computer.CPU != "AMD Ryzen 9" {
		t.Errorf("Expected CPU to be 'AMD Ryzen 9', but got '%s'", computer.CPU)
	}

	if computer.GPU != "NVIDIA RTX 3090" {
		t.Errorf("Expected GPU to be 'NVIDIA RTX 3090', but got '%s'", computer.GPU)
	}

	if computer.RAM != 64 {
		t.Errorf("Expected RAM to be 64, but got %d", computer.RAM)
	}

	if computer.Storage != 2048 {
		t.Errorf("Expected Storage to be 2048, but got %d", computer.Storage)
	}

	if computer.Display != "5K" {
		t.Errorf("Expected Display to be '5K', but got '%s'", computer.Display)
	}
}
