package main

import (
	"fmt"
	"strconv"
)

func IsPalindrome(num int) bool {
	str := strconv.Itoa(num)
	
	for i, j := 0, len(str) - 1; i <= j; i, j = i + 1, j - 1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func main() {
	max := 0

	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			if IsPalindrome(j * i) && j*i > max {
				max = j*i
			}
		}
	}

	fmt.Println(max)
}