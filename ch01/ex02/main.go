package main

import (
	"os"
	"strconv"
)

func main() {
	for index, arg := range os.Args[1:] {
		indexString := strconv.Itoa(index)
		println(indexString + ": " + arg)
	}
}
