package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./distconv"
)

func show(t float64) {
	mile := distconv.Mile(t)
	meter := distconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n", mile, distconv.MiToM(mile), meter, distconv.MToMi(meter))
}

func toF16(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	return num
}

func main() {
	if len(os.Args[1:]) == 0 {
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			num := toF16(stdin.Text())
			show(num)
		}
	} else {
		for _, arg := range os.Args[1:] {
			num := toF16(arg)
			show(num)
		}
	}
}
