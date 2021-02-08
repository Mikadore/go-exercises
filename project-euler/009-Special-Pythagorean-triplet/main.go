package main

import (
	"fmt"
	"math"
)

func main() {
	for b := 1;; b++ {
		for a := 1; a < b; a++ {
			c := math.Sqrt(float64(a*a + b*b))
			if float64(a) + float64(b) + c == 1000.0 {
				fmt.Printf("%08f", float64(a)*float64(b)*c)
				return
			}
		}
	}
}