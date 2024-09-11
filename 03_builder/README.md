

**建造者模式（Builder Pattern）**是一种创建型设计模式，目的是通过将复杂对象的构建过程分离出来，使客户端只需要关心建造的过程，而无需了解对象内部的组成细节。在 Go 中，建造者模式可以用于创建复杂的对象，同时可以提供链式方法调用来简化对象的构建过程。

## 1. 需求：创建一台复杂的计算机

我们可以创建一个 `Computer` 对象，它包含多个组件（CPU、GPU、RAM、存储等）。使用建造者模式，可以让用户通过链式调用构建 `Computer` 对象，而不需要关心每个组件的详细构造过程。

## 2.建造者模式流程图（基于计算机构建）

```markdown
+----------------------------------+
|             Director             |
|       (main function)            |
+----------------------------------+
                |
                v
+----------------------------------+
|        ComputerBuilder           |
|      (构建 Computer 对象)         |
+----------------------------------+
| - SetCPU(cpu string)             |
| - SetGPU(gpu string)             |
| - SetRAM(ram int)                |
| - SetStorage(storage int)        |
| - SetDisplay(display string)     |
| - Build()                        |
+----------------------------------+
                |
                v
+----------------------------------+
|           Computer               |
| (最终构建的复杂对象)              |
+----------------------------------+
| - CPU: Intel i7                  |
| - GPU: NVIDIA RTX 3080           |
| - RAM: 32GB                      |
| - Storage: 1024GB                |
| - Display: 4K                    |
+----------------------------------+
```

流程描述：

1. **Director（指挥者）**：
   - 在该例子中，`main` 函数充当指挥者角色，负责指导整个构建流程。它调用 `ComputerBuilder` 的各种设置方法，最终生成一个 `Computer` 对象。

2. **Builder（建造者）**：
   - `ComputerBuilder` 是具体的建造者，实现了多个构建步骤，例如设置 CPU、GPU、RAM 等。每个方法返回构建器自身，以实现链式调用。

3. **Product（产品）**：
   - `Computer` 是最终构建的产品对象，包含了 CPU、GPU、RAM 等具体的配置。

关键步骤：

1. **初始化 Builder**：通过 `NewComputerBuilder()` 方法创建一个 `ComputerBuilder` 实例。
2. **设置属性**：通过链式调用，逐步为 `ComputerBuilder` 设置 CPU、GPU、RAM、存储和显示器等属性。
3. **构建最终对象**：调用 `Build()` 方法，最终生成 `Computer` 对象，该对象包含之前设置的所有配置。

## 3. 建造者模式实现

示例代码

```go
package main

import "fmt"

// Computer 是我们要构建的复杂对象
type Computer struct {
	CPU      string
	GPU      string
	RAM      int
	Storage  int
	Display  string
}

// ComputerBuilder 是建造者，用于一步步构建 Computer 对象
type ComputerBuilder struct {
	cpu      string
	gpu      string
	ram      int
	storage  int
	display  string
}

// NewComputerBuilder 创建一个新的 ComputerBuilder
func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

// SetCPU 设置 CPU
func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.cpu = cpu
	return b
}

// SetGPU 设置 GPU
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
	// 使用建造者模式构建计算机
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
```

## 4. 示例输出

```bash
Computer Config: CPU=Intel i7, GPU=NVIDIA RTX 3080, RAM=32GB, Storage=1024GB, Display=4K
```

## 5. 单元测试

我们可以通过 Go 的 `testing` 包为这个建造者模式添加单元测试，确保其正常工作。

单元测试代码

```go
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
```

## 6. 总结

- **建造者模式**适合用来构建复杂的对象，特别是当对象有很多可选参数时，通过链式调用来设置这些参数可以使代码更加简洁易读。
- 在实际项目中，建造者模式广泛应用于对象构建、配置类、客户端设置等场景，可以有效避免构造函数参数过多导致的代码可读性问题。