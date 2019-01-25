package str

import "testing"

func TestStringRand(t *testing.T) {
	r := Rand(10)

	t.Log(r)
	if len(r) != 10 {
		t.Error("Should return string with 10 length")
	}
}

func TestUpperFist(t *testing.T) {

	s := "abc"

	r := UpperFirst(s)

	if r != "Abc" {
		t.Error("abc should upper first to Abc")
	}
}