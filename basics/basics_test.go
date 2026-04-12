package basics

import (
	"testing"
)

func TestString(t *testing.T) {
	s := "hello"
	t.Log("something log")
	if s != "hello" {
		t.Fatalf("got %q", s)
	}
}
