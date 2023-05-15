package Factory

import "testing"

func TestNewPeople(t *testing.T) {
	NewPeople("t").GetName()
	NewPeople("s").GetName()
}
