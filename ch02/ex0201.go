package main

import "fmt"

func main() {
	// x := complex(2.5, 3.1)
	// y := complex(10.2, 2)
	// fmt.Println(x + y)
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(z, d)
}
