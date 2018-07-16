package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 || y&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	var f float64 = 3.13
	var fMinus float64 = -3.13
	fmt.Println(int(f))
	fmt.Println(int(fMinus))
}
