package main

import (
	"fmt"

	"matrix"
)

func main() {
	m := matrix.NewZeroMatrix[int](5, 5)
	fmt.Println(m)
}
