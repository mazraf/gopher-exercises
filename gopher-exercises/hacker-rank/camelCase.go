package main

import (
	"fmt"
)

func rank() {
	var input string
	fmt.Scanf("%s\n", &input)
	min, max := 'A', 'Z'
	ans := 1
	for _, ch := range input {
		if ch >= min && ch <= max {
			ans++
		}
	}
	fmt.Println(ans)
}
