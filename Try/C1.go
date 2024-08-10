package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func abc() {
	fmt.Printf("Type: %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v\n", z, z)
}

func dtype() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func main() {
	// abc()
	// dtype()
	// sqrtConvertiion()
	// fmt.Println(needInt(Small))

}

func C1() {
	var i, j int = 1, 2
	k := 3

	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

// Datatype Convertiion

func sqrtConvertiion() {
	var x, y int = 7, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}
