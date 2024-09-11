# 工厂模式(Factory)

详情可以查看博客: https://lailin.xyz/post/factory.html

## 简单工厂

由于 Go 本身是没有构造函数的，一般而言我们采用 NewName  的方式创建对象/接口，当它返回的是接口的时候，其实就是简单工厂模式

## 工厂方法

当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，我们推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂

## 抽象工厂

一个工厂方法可以创建**相关联**的多个类的时候就是抽象工厂模式，这个不太常用

## DI 容器
我们这里的实现比较粗糙，但是作为一个 demo 理解 di 容器也足够了，和 dig 相比还缺少很多东西，并且有许多的问题，例如 依赖关系，一种类型如果有多个 provider 如何处理等等等等。

---



在 Go 语言中，**工厂模式**是一种用于创建对象的设计模式，目的是将实例化对象的逻辑与具体使用对象的逻辑分离，特别适合当创建过程较复杂时。工厂模式可以避免客户端直接创建对象，而是通过工厂方法来获取实例。

下面展示如何在 Go 中实现工厂模式。

### 1. 简单工厂模式

简单工厂模式通过一个工厂类（或方法）来创建不同的对象。客户端只需要传递类型或条件，不需要知道对象的具体实现细节。

#### 示例代码

```go
package main

import "fmt"

// Product 是一个接口，所有产品类型都实现这个接口
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
		fmt.Println(productA.GetName()) // 输出：Product A
	}

	// 使用工厂创建产品 B
	productB := CreateProduct("B")
	if productB != nil {
		fmt.Println(productB.GetName()) // 输出：Product B
	}
}
```

#### 说明：
- **Product 接口**：定义了所有产品需要实现的公共接口。
- **ConcreteProductA 和 ConcreteProductB**：具体的产品实现，分别代表不同的产品类型。
- **CreateProduct 函数**：工厂方法，根据传入的 `productType` 返回对应的产品对象。

### 2. 工厂方法模式

工厂方法模式是工厂模式的进一步抽象。在这种模式中，创建对象的任务被委托给子类。每个子类都有自己的工厂方法，负责创建特定类型的对象。

#### 示例代码

```go
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
```

#### 说明：
- **Product 接口**：同样定义了所有产品需要实现的公共接口。
- **ConcreteFactoryA 和 ConcreteFactoryB**：具体工厂，负责创建不同的产品。
- **Factory 接口**：定义工厂类需要实现的接口方法 `CreateProduct`。

### 3. 抽象工厂模式

抽象工厂模式用于创建一系列相关或互相依赖的对象，而无需指定它们的具体类。适用于需要创建多个相关对象的场景。

#### 示例代码

```go
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
	fmt.Println(adidasShoe.GetShoeType()) // 输出：Adidas Shoe
	fmt.Println(adidasShirt.GetShirtType()) // 输出：Adidas Shirt

	// Nike 工厂
	nikeFactory := &NikeFactory{}
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()
	fmt.Println(nikeShoe.GetShoeType())   // 输出：Nike Shoe
	fmt.Println(nikeShirt.GetShirtType()) // 输出：Nike Shirt
}
```

#### 说明：
- **Shoe 和 Shirt 接口**：代表产品族中的不同产品类型（如鞋子和衣服）。
- **Adidas 和 Nike 的产品实现**：分别实现 `Shoe` 和 `Shirt`，代表不同品牌的产品。
- **AdidasFactory 和 NikeFactory**：具体工厂，负责创建一系列相关的产品（鞋子和衣服）。

### 总结

- **简单工厂模式**：通过一个工厂函数，根据条件返回不同的产品实例，适合简单场景。
- **工厂方法模式**：通过不同的工厂类，创建不同类型的产品，适合扩展性强的场景。
- **抽象工厂模式**：用于创建相关联的一系列产品，适合复杂场景下的多种产品族的生产。