package main

import (
	"fmt"
)

func main() {
	x := 2

	fmt.Printf("%d\t%b\n", x, x)

	y := 2 << 1

	fmt.Printf("%d\t%b", y, y)
}
