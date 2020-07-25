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

// Struct
type Vertex struct {
	X int
	Y int
}

func main() {
	// Pointer - no pointer arithmetic
	pointer()

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
