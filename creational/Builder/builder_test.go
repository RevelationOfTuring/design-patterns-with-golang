package Builder

import "testing"

func TestDirect_Construct(t *testing.T) {
	builder := NewConcreteBuilder()
	d := NewDirector(builder)
	d.Construct()
}
