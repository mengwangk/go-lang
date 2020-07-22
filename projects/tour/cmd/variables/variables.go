package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func add_v2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Variable declarations
var c, python, java bool
var i, j int = 1, 2

// Numeric constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Now you have %g problems.\n", math.Pow(7, 2))

	fmt.Println("PI value is ", math.Pi)

	fmt.Println(add(3, 2))
	fmt.Println(add_v2(1, 2))

	fmt.Println(swap("hello", "world"))

	// Naked return
	fmt.Println(split(171))

	var i int
	fmt.Println(i, c, python, java)

	var a, b, c = true, false, "no!"
	fmt.Println(i, j, a, b, c)

	k := 3
	fmt.Println(k, i, j, a, b, c)

	// Var block
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Default values
	var x int
	var f float64
	var y bool
	var s string
	fmt.Printf("%v %v %v %q\n", x, f, y, s)

	// Type conversion
	m, n := 3, 4
	x, y = 4, true
	var p float64 = math.Sqrt(float64(m*m + n*n))
	var q uint = uint(p)
	fmt.Println(m, n, q)

	var d int = 42
	var e float64 = float64(i)
	var g uint = uint(f)
	fmt.Printf("%g %v %v\n", d, e, g)

	// Type inference
	v := 42.1 // change me!
	fmt.Printf("v is of type %T\n", v)

	// Constant
	const Pi = 3.14
	fmt.Println(Pi)

	// Numeric constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
