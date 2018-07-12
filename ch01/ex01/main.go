package main

import (
	"fmt"
	"os"
)

func main() {
	var result, separator string
	for i := 0; i < len(os.Args); i++ {
		result += separator + os.Args[i]
		separator = " "
	}
	fmt.Println(result)
}
