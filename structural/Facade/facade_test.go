package Facade

import "testing"

func TestNewPhoneFacade(t *testing.T) {
	f := NewPhoneFacade()
	f.setup()
}
