package main

import (
	"fmt"
)

const (
	rue = 0 == 0
)

func main() {
	foo := []string{"A"}
	bar := [2]string{"V", "W"}
	fmt.Println(foo, append(foo, "B"))
	fmt.Printf("type of foo is %T\n", foo)
	fmt.Printf("capability of foo is %v\n", cap(foo))
	fmt.Printf("capability of bar is %v\n", cap(bar))

	foo = append(foo, "B", "C", "D")
	fmt.Printf("capability of foo is %v after append\n", cap(foo))
	fmt.Printf("lenth of foo is %v\n", len(foo))

	fmt.Println("Copy slice----------------------------------------")

	s := []string{"asm", "d", "e"}
	d := []string{"b"}
	n := copy(d, s)
	fmt.Println("copyed slice is", d, ", size of copied: ", n)
	fmt.Println("--------------------------------------------------")

	fmt.Println("Delete element in map-----------------------------")
	m := map[string]string{
		"tom": "A",
		"cat": "B",
	}
	delete(m, "tom")
	fmt.Println("map is ", m)
	fmt.Println("--------------------------------------------------")

	fmt.Println("len-----------------------------------------------")
	fmt.Println("len of array is: ", len([3]int{1, 2, 3}))
	fmt.Println("len of array pointer is ", len(&([2]int{})))
	fmt.Println("len of string is ", len("ABCDEF"))
	fmt.Println("--------------------------------------------------")

	fmt.Println("make----------------------------------------------")
	s1 := make([]int, 3)
	fmt.Println("make slice", s1, "capability is", cap(s1), "length is", len(s1))
	s2 := make([]int, 2, 4)
	fmt.Println("make slice", s2, "capability is", cap(s2), "length is", len(s2))
	fmt.Println("make map", make(map[string]int))
	fmt.Println("--------------------------------------------------")

	fmt.Println("new ----------------------------------------------")
	fmt.Println("new type", new(string))
	fmt.Println("new expression ", new(2))
	fmt.Println("--------------------------------------------------")
}
