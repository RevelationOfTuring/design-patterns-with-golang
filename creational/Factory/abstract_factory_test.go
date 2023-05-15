package Factory

import "testing"

func TestNewFactory(t *testing.T) {
	factory := NewFactory()
	food := factory.makeFood()
	food.GetPrice()
	drink := factory.makeDrink()
	drink.GetPrice()
}
