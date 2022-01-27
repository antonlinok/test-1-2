package main

import "testing"

func Testitoa(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a", "a"},
		{"123", "123"},
		{"-123", "-123"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := itoa(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
