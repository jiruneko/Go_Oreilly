package main

import (
	"fmt"
	"math"
)

func main() {
	// var x = [3]int{10, 20, 30}
	// var y = []int{10, 20, 30}
	// fmt.Println(x)
	// fmt.Println(y)
	fmt.Println((math.Pi))
	// a := []int{1, 2, 3, 4}
	// b := a[:2]
	// fmt.Println(cap(a), cap(b)) // 4 4
	// b = append(b, 30)
	// fmt.Println("a:", a) // x: [1 2 30 4]
	// fmt.Println("b:", b)
	x := make([]int, 0, 5)              // []
	x = append(x, 1, 2, 3, 4)           // [1 2 3 4]
	y := x[:2]                          // [1 2]
	z := x[2:]                          // [3 4]
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	y = append(y, 30, 40, 50)           // [1 2 30 40 50]
	x = append(x, 60)                   // [1 2 30 40 60]
	z = append(z, 70)                   // [30 40 70]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
