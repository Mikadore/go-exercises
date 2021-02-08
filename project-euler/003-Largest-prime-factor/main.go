package main

import (
	"fmt"
)
func factor(num uint) uint {
	for i := uint(2); i < num/2; i++ {
		if num % i == 0 {
			return i
		}
	}
	return num
}
func main() {
	var num 	uint = 600851475143
	
	for {
		nfact := factor(num)

		num /= nfact

		fmt.Println(nfact)

		if num == 1 {
			break
		}

	}
}