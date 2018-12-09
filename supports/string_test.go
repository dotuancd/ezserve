package supports

import "testing"

func TestStringRand(t *testing.T) {
	r := StringRand(10)

	t.Log(r)
	if len(r) != 10 {
		t.Error("Should return string with 10 length")
	}
}