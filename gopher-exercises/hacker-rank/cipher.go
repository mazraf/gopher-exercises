package main

import (
	"fmt"
)

func main() {
	var input string
	var l, delta int
	fmt.Scanf("%d\n", &l)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)
	var res []rune
	for _, r := range []rune(input) {
		res = append(res, cipher(r, delta))
	}
	fmt.Println(string(res))
}
func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	} else if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	} else {
		return r
	}
}
func rotate(r rune, base, delta int) rune {
	temp := int(r) - base
	temp = (temp + delta) % 26
	return rune(temp + base)
}
