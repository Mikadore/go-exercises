/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main 

import "fmt"

func gcd(a,b uint) uint {
	if b == 0 {
		return a
	} else {
		return gcd(b, a % b)
	}
}

func lcm(a,b uint) uint {
	return a*b / gcd(a, b);
}

func list_lcm(arr []uint) (uint, bool) {
	if len(arr) < 2 {
		return 0, false
	}
	x := lcm(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		x = lcm(x, arr[i])
	}

	return x, true
}

func main() {
	//easier to work out by hand
	//	fmt.Println(2*2*2*2*3*3*5*7*11*13*17*19)
	fmt.Print(list_lcm([]uint{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}))

}