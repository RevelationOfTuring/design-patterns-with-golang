package Factory

import "fmt"

type Lunch interface {
	GetPrice()
}

type Rice struct{}

func (r *Rice) GetPrice() {
	fmt.Println("大米一元")
}

type Wine struct{}

func (w *Wine) GetPrice() {
	fmt.Println("大米二元")
}

type LunchFactory interface {
	makeFood() Lunch
	makeDrink() Lunch
}

// 通过不同的LunchFactory实现类，来生产不同的Lunch实现类
type SimpleFactory struct{}

func (sf *SimpleFactory) makeFood() Lunch {
	return &Rice{}
}

func (sf *SimpleFactory) makeDrink() Lunch {
	return &Wine{}
}

func NewFactory() LunchFactory {
	return &SimpleFactory{}
}
