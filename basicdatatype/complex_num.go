package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x, y)                      // 1+2i 3+4i
	fmt.Println(x*y, real(x*y), imag(x*y)) // -5+10i -5 10

	a := 1 + 2i
	b := 3 + 4i
	fmt.Println(a, b)

	fmt.Println(cmplx.Sqrt(-1))
}
