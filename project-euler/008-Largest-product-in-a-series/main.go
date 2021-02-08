package main

import (
	"fmt"
	"io/ioutil"
)
func main() {

	dat, err := ioutil.ReadFile("numbah.txt")

	if err != nil {
		panic(err)
	}

	max := 0
	//fmt.Println(dat)
	for i := 0; i + 13 < len(dat) - 1; {
		mul := 1
		for j := i; j < i + 13; j++ {
			mul *= int(dat[j] - 48)
		}
		if mul > max {
			max = mul
		}

		if dat[i + 13] != '0' {
			i++
		} else {
			i += 14
		}
	}

	fmt.Println(max)


}