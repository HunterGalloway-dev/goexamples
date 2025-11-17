package simplemath_test

import(
	"simpletest/simplemath"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    int
		y    int
		want int
	}{
		{
			name: "Add two positive numbers",
			x: 5,
			y: 9,
			want: 14,
		},
		{
			name: "Add two negative numbers",
			x: -5,
			y: -9,
			want: -14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simplemath.Add(tt.x, tt.y)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

