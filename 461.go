package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := hammingDistance(1, 4)
	fmt.Println(b)
}

func hammingDistance(x int, y int) int {
	s1 := strconv.FormatInt(int64(x), 2)
	s2 := strconv.FormatInt(int64(y), 2)

	// var u string
	// if len(s1) > len(s2) {
	// 	for i := 0; i < len(s1)-len(s2); i++ {
	// 		u = u + "0"
	// 	}
	// 	s2 = u + s2
	// } else {
	// 	for i := 0; i < len(s2)-len(s1); i++ {
	// 		u = u + "0"
	// 	}
	// 	s1 = u + s1
	// }

	fmt.Println(s1)
	fmt.Println(s2)
	a := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			a++
		}
	}
	return a
}
