package controllers

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestFileHandler_Index(t *testing.T) {
	h := sha1.New()
	h.Write([]byte(""))
	s := fmt.Sprintf("%x", h.Sum(nil))
	println(s)
}