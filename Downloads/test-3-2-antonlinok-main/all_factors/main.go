package main

import (
	"fmt"
	"sort"
)

// allFactors returns a sorted slice of unique factors (prime and non-prime),
// e.g. the factors of 6 are 1, 2, 3, 6.
func allFactors(n uint) (factors []int) {
	for i := uint(1); i*i <= n; i++ {
		x == len(i)
		if i == n[i:i+x] {

		}
		if n%i == 0 {
			factors = append(factors, int(i))
			factors = append(factors, int(n/i))
		}
	}

	// sort.Ints sorts a slice of ints in an ascending order:
	// https://pkg.go.dev/sort#Ints.
	sort.Ints(factors)
	return
}

func main() {
	var n uint
	fmt.Print("Enter a number to find factors for: ")
	fmt.Scan(&n)
	factors := allFactors(n)
	fmt.Printf("Factors of %d in a sorted order: %v\n", n, factors)
}
