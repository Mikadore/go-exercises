package main

import "fmt"



func main() {

	numbers := make([]bool, 2000000)

	for i := range numbers {
		numbers[i] = true
	}

	for i := 2; i < len(numbers); {
		for j := 2*i; j < len(numbers); j += i {
			numbers[j] = false
		}
		for i++;i < len(numbers) && !numbers[i];i++ {}
	}

	sum := 0

	for i, v := range numbers {
		if v {
			sum += i
		}
	}
	sum--
	fmt.Println(sum)



}