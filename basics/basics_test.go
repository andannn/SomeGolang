package basics

import (
	"math"
	"somegolang/internal/testutil"
	"testing"
)

func TestSwap(t *testing.T) {
	f := func(x, y string) (string, string) {
		return y, x
	}

	a, b := f("hello", "world")
	testutil.AssertEqual(t, a, "world")
	testutil.AssertEqual(t, b, "hello")
}

func TestGoPassByValue_String(t *testing.T) {
	change := func(s string) {
		s = "change this string"
	}

	original := "hello"

	change(original)

	testutil.AssertEqual(t, original, "hello")
}

func TestNamedReturnValues(t *testing.T) {
	split := func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}

	a, b := split(17)
	testutil.AssertEqual(t, a, 7)
	testutil.AssertEqual(t, b, 10)
}

func TestForLoop1(t *testing.T) {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += i
	}

	testutil.AssertEqual(t, sum, 3)
}

func TestForLoop2(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	testutil.AssertEqual(t, sum, 1024)
}

func TestForLoop3(t *testing.T) {
	sum := 1
	for {
		sum++
		if sum == 3 {
			break
		}
	}

	testutil.AssertEqual(t, sum, 3)
}

func TestForLoop4(t *testing.T) {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		t.Log("index", i, "value", v)
	}
}

func TestDefer1(t *testing.T) {
	f := func() {
		i := 0
		// defer会把函数调用压到 defer 栈里，同时当场把这次调用要用到的参数先算出来。
		defer testutil.AssertEqual(t, i, 0)
		i++
	}

	f()
}

func TestDefer2(t *testing.T) {
	f := func() {
		i := 0
		// 用闭包捕获i，defer 执行时找到i的值
		defer func() {
			testutil.AssertEqual(t, i, 1)
		}()

		i++
	}

	f()
}

func TestSturct(t *testing.T) {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 3}
	testutil.AssertEqual(t, v.X, 1)
	testutil.AssertEqual(t, v.Y, 3)
}

func TestStructWithPointer(t *testing.T) {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 3}
	p := &v
	p.X = 2
	testutil.AssertEqual(t, v.X, 2)
	testutil.AssertEqual(t, v.Y, 3)
}

type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func TestStructWithMethod(t *testing.T) {
	f := MyFloat(-1)
	testutil.AssertEqual(t, f.abs(), 1)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestStructWithMethod2(t *testing.T) {
	v := Vertex{3, 4}
	//	v.Scale(10)
	(&v).Scale(10)
	testutil.AssertEqual(t, v.X, 30)
	testutil.AssertEqual(t, v.Y, 40)
}

func TestStructWithMethod3(t *testing.T) {
	v := Vertex{3, 4}
	// abs := v.Abs()
	p := &v
	abs := (*p).Abs()
	testutil.AssertEqual(t, abs, 5)
}

