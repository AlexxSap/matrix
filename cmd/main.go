package main

import (
	"fmt"

	"matrix"
)

func main() {
	m := matrix.NewMatrix[int](5, 5)
	fmt.Println(m)
}
