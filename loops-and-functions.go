package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z1 := float64(1)
	z2 := z1 - (z1*z1-x)/(2*z1)
	count := 1

	for (z1-z2) < -0.000001 || (z1-z2) > 0.000001 {
		z1, z2 = z2, z1
		z2 = z1 - (z1*z1-x)/(2*z1)
		count++
	}
	fmt.Println(count)
	return z2
}

func main() {
	fmt.Println(Sqrt(2))
}
