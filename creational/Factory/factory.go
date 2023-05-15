package Factory

import "fmt"

type People interface {
	GetName()
}

type Student struct{}

func (s *Student) GetName() {
	fmt.Println("学生的名字")
}

type Teacher struct{}

func (t *Teacher) GetName() {
	fmt.Println("老师的名字")
}

func NewPeople(name string) People {
	// name可以看做是工厂的标识符
	switch name {
	case "s":
		return &Student{}
	case "t":
		return &Teacher{}
	default:
		return nil
	}
}
