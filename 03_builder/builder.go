package main

import "fmt"

// Computer 是我们要构建的复杂对象
type Computer struct {
	CPU     string
	GPU     string
	RAM     int
	Storage int
	Display string
}

// ComputerBuilder 是建造者，用于一步步构建 Computer 对象
type ComputerBuilder struct {
	cpu     string
	gpu     string
	ram     int
	storage int
	display string
}

// NewComputerBuilder 创建一个新的 ComputerBuilder
func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

// SetCPU 设置CPU
func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.cpu = cpu
	return b
}

// SetGPU 设置GPU
func (b *ComputerBuilder) SetGPU(gpu string) *ComputerBuilder {
	b.gpu = gpu
	return b
}

// SetRAM 设置 RAM 大小
func (b *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
	b.ram = ram
	return b
}

// SetStorage 设置存储大小
func (b *ComputerBuilder) SetStorage(storage int) *ComputerBuilder {
	b.storage = storage
	return b
}

// SetDisplay 设置显示器类型
func (b *ComputerBuilder) SetDisplay(display string) *ComputerBuilder {
	b.display = display
	return b
}

// Build 构建最终的 Computer 对象
func (b *ComputerBuilder) Build() *Computer {
	return &Computer{
		CPU:     b.cpu,
		GPU:     b.gpu,
		RAM:     b.ram,
		Storage: b.storage,
		Display: b.display,
	}
}

func main() {
	builder := NewComputerBuilder()
	computer := builder.SetCPU("Intel i7").
		SetGPU("NVIDIA RTX 3080").
		SetRAM(32).
		SetStorage(1024).
		SetDisplay("4K").
		Build()
	// 打印计算机配置
	fmt.Printf("Computer Config: CPU=%s, GPU=%s, RAM=%dGB, Storage=%dGB, Display=%s\n",
		computer.CPU, computer.GPU, computer.RAM, computer.Storage, computer.Display)
}
