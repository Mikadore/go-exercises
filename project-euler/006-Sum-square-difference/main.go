package main

import "fmt"

func main() {
	const k = 100

	sum_squares := 0

	sum := 0
	for i := 1; i <= k; i++ {
		sum_squares += i*i
		sum += i
	}

	sum_squared := sum*sum

	fmt.Println(sum_squared - sum_squares)

}