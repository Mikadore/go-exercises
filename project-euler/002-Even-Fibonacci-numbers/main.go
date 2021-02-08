package main 

import "fmt"

func fib() func() uint {
	var a, b uint = 1, 1
	return func() (x uint) {
		x, a, b = a, b, a + b
		return
	}
}

func main() {
	var f = fib()

	sum := uint(0)
	
	for a, b := uint(0), uint(0); a + b + b < 4000000; {
		a, b = f(), f()

		sum += f()
	}

	fmt.Println(sum)
}