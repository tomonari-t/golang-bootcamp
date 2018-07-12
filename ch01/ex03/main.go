package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func showCliArgs() {
	var result, separator string
	for _, arg := range os.Args[1:] {
		result += separator + arg
		separator = " "
	}
	fmt.Println(result)
}

func showCliArgsWithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start1 := time.Now()
	showCliArgs()
	duration1 := time.Until(start1)

	start2 := time.Now()
	showCliArgsWithJoin()
	duration2 := time.Until(start2)

	fmt.Println(duration1.Nanoseconds())
	fmt.Println(duration2.Nanoseconds())
}
