package main

import (
	"bufio"
	"somegolang/internal/testutil"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	r := strings.NewReader("hello\nworld")
	bufr := bufio.NewReader(r)
	peeked, _ := bufr.Peek(2)
	testutil.AssertEqual(t, string(peeked), "he")
}
