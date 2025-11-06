package TwoSum

import (
	"testing"
)

/*
Test strategy:
- Basic cases with different array sizes
- Negative numbers (should work fine)
- Zero in the mix
- Answer at different positions (start, middle, end)

Go's testing is simple - just check if the result matches expected
No fancy assertions needed, just compare slices
*/

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic case from problem",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "answer at the end",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "same number twice",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "with zero",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			// Check if we got a result
			if got == nil {
				t.Errorf("twoSum() returned nil, want %v", tt.want)
				return
			}

			// Check if indices match (order doesn't matter per problem spec)
			if !equalPairs(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper to compare index pairs - order doesn't matter
func equalPairs(a, b []int) bool {
	if len(a) != 2 || len(b) != 2 {
		return false
	}
	return (a[0] == b[0] && a[1] == b[1]) || (a[0] == b[1] && a[1] == b[0])
}
