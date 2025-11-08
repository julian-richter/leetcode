package main_test

import (
	main "longest-substring/cmd"
	"testing"
)

/*
Test strategy:
- Empty string (edge case)
- All unique chars
- All the same chars (the worst case)
- Mix of duplicates
- Single character
- Long strings with patterns

table-driven tests are the Go way - clean and easy to add more cases
*/

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example from problem",
			s:    "abcabcbb",
			want: 3, // "abc"
		},
		{
			name: "all same character",
			s:    "bbbbb",
			want: 1, // "b"
		},
		{
			name: "all unique",
			s:    "pwwkew",
			want: 3, // "wke"
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "two unique chars",
			s:    "au",
			want: 2,
		},
		{
			name: "long unique string",
			s:    "abcdefghijklmnopqrstuvwxyz",
			want: 26,
		},
		{
			name: "duplicate at start",
			s:    "aab",
			want: 2, // "ab"
		},
		{
			name: "string with space all unique",
			s:    "ab cd",
			want: 5, // entire string is unique
		},
		{
			name: "repeating pattern",
			s:    "dvdf",
			want: 3, // "vdf"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.LengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
