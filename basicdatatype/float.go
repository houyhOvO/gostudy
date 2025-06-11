package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216
	fmt.Println(f == f+1) // true !!!

	// Prints the power of e with three decimal places and 8 characters.
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN

	nan := math.NaN()
	fmt.Println(nan == nan, nan > nan, nan < nan) // false false false
}
