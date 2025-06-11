package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255, 0 (overflow), 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127 -128 1

	var x uint8 = 1<<1 | 1<<5 // x<<n equals to multiply 2^n, x>>n equals to be divided by 2^n
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x) // 00100010
	fmt.Printf("%08b\n", y) // 00000110

	fmt.Printf("%08b\n", x&y)  // 00000010
	fmt.Printf("%08b\n", x|y)  // 00100110
	fmt.Printf("%08b\n", x^y)  // 00100100
	fmt.Printf("%08b\n", x&^y) // ^y = 11111001 x&^y = 00100000

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) // 1 and 5
		}
	}

	f := 3.141
	g := int(f)
	fmt.Println(f, g) // 3.141 3
}
