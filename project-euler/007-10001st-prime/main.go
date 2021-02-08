package main

import "fmt"	

func IsPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {

	counter, i := 0, 2

	for ;;i++ {
		if IsPrime(i){
			counter++
			if counter == 10001 {
				fmt.Println(i)
				break;
			}
		}
	}

}