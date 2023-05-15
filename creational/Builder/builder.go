package Builder

import "fmt"

// 将生产过程解耦开，分成不同部分

type Builder interface {
	FirstDo() string
	SecondDo() string
	ThirdDo() string
	FinallyDo() string
}

// 外部直接操作的对象
type Director struct {
	b Builder
}

func NewDirector(b Builder) *Director {
	return &Director{b}
}

// 生产过程（调用builder中的各步骤并结合）
func (d *Director) Construct() {
	// 生产产品
	res := fmt.Sprintf(`1. %s
2. %s
3. %s
finally. %s`,
		d.b.FirstDo(), d.b.SecondDo(), d.b.ThirdDo(), d.b.FinallyDo())
	fmt.Println(res)
}

// 具体builder的实现类
var _ Builder = &ConcreteBuilder{} // 用于goland实现类的各实现方法补全

type ConcreteBuilder struct{}

func NewConcreteBuilder() Builder {
	return &ConcreteBuilder{}
}

func (c *ConcreteBuilder) FirstDo() string {
	return "first done"
}

func (c *ConcreteBuilder) SecondDo() string {
	return "second done"
}

func (c *ConcreteBuilder) ThirdDo() string {
	return "third done"
}

func (c ConcreteBuilder) FinallyDo() string {
	return "final done"
}
