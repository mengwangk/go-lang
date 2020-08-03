package main

import "fmt"

// Pointer
func pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func pointer2() {
	// taking a normal variable
	var x int = 5748

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &x

	// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
	fmt.Println("Deferenced value in variable p = ", *p)
}

// Struct
type Vertex struct {
	X int
	Y int
}

func pointerToStruct() {
	// Struct
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	// Pointer to struct
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func structLiteral() {
	fmt.Println(v1, p, v2, v3)
}

func array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func main() {
	pointer()

	pointer2()

	pointerToStruct()

	structLiteral()

	array()

	slice()

}
