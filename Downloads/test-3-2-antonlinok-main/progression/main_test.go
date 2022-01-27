package main

import (
	"testing"
)

func TestnthTerm(t *testing.T) {
    for _, tc := range []struct{
        s    n uint
        want int
    }{
        {uint{}, []int{}},
        {uint{1}, []int{3}},
        {uint{4}, []int{-3}}, 
        {uint{0}, []int{}},
    } {
        got := nthTerm(tc.n)
        if !equal(got, tc.want) {
            t.Errorf("nthTerm(%v) = %v, want = %v", tc.n, got, tc.want)
        }

    }
}

func equal(a, b int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}