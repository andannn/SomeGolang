package main

import (
	"bufio"
	"bytes"
	"somegolang/internal/testutil"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	r := strings.NewReader("hello\nworl的")
	bufr := bufio.NewReader(r)
	peeked, _ := bufr.Peek(2)
	testutil.AssertEqual(t, string(peeked), "he")

	b := make([]byte, 2)
	bytesRead, _ := bufr.Read(b)
	testutil.AssertEqual(t, 2, bytesRead)
	testutil.AssertEqual(t, string(b), "he")

	stringRead, _ := bufr.ReadString('\n')
	testutil.AssertEqual(t, stringRead, "llo\n")

	sliceRead, _ := bufr.ReadSlice('l')
	testutil.AssertEqual(t, string(sliceRead), "worl")

	testutil.AssertEqual(t, bufr.Size(), 4096)

	runeRead, n, _ := bufr.ReadRune()
	testutil.AssertEqual(t, n, 3)
	testutil.AssertEqual(t, string(runeRead), "的")
}

func TestWriter(t *testing.T) {
	var buf bytes.Buffer

	w := bufio.NewWriter(&buf)
	w.WriteString("hello")
	testutil.AssertEqual(t, buf.String(), "")

	testutil.AssertEqual(t, w.Buffered(), 5)

	w.Flush()
	testutil.AssertEqual(t, buf.String(), "hello")

	avalible := w.Available()
	testutil.AssertEqual(t, avalible, 4096)

	b := make([]byte, 4097, 4097)
	w.Write(b)

	w.Flush()

	testutil.AssertEqual(t, buf.Len(), 4102)
}
