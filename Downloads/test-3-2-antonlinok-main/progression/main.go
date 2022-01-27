package main

import "fmt"

func nthTerm(n uint) int {
	a, b, c := 5, 1, 3
	var i uint
	for ; i < n; i = i + 2 {
		a, b = b, a-2*b*1
		a, c = c, -2*a*1
	}
	return a
}
func main() {
	var n uint
	fmt.Println("Enter n:")
	fmt.Scan(&n)
	fmt.Println(nthTerm(n))
}
