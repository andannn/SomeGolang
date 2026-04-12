package basics

import (
	"math"
	"testing"
)

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSwap(t *testing.T) {
	f := func(x, y string) (string, string) {
		return y, x
	}

	a, b := f("hello", "world")
	assertEqual(t, a, "world")
	assertEqual(t, b, "hello")

}

func TestGoPassByValue_String(t *testing.T) {
	change := func(s string) {
		s = "change this string"
	}

	original := "hello"

	change(original)

	assertEqual(t, original, "hello")
}

func TestNamedReturnValues(t *testing.T) {
	split := func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}

	a, b := split(17)
	assertEqual(t, a, 7)
	assertEqual(t, b, 10)
}

func TestForLoop1(t *testing.T) {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += i
	}

	assertEqual(t, sum, 3)
}

func TestForLoop2(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	assertEqual(t, sum, 1024)
}

func TestForLoop3(t *testing.T) {
	sum := 1
	for {
		sum++
		if sum == 3 {
			break
		}
	}

	assertEqual(t, sum, 3)
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
		defer assertEqual(t, i, 0)
		i++
	}

	f()
}

func TestDefer2(t *testing.T) {
	f := func() {
		i := 0
		// 用闭包捕获i，defer 执行时找到i的值
		defer func() {
			assertEqual(t, i, 1)
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
	assertEqual(t, v.X, 1)
	assertEqual(t, v.Y, 3)
}

func TestStructWithPointer(t *testing.T) {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 3}
	p := &v
	p.X = 2
	assertEqual(t, v.X, 2)
	assertEqual(t, v.Y, 3)
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
	assertEqual(t, f.abs(), 1)
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
	assertEqual(t, v.X, 30)
	assertEqual(t, v.Y, 40)
}

func TestStructWithMethod3(t *testing.T) {
	v := Vertex{3, 4}
	// abs := v.Abs()
	p := &v
	abs := (*p).Abs()
	assertEqual(t, abs, 5)
}
