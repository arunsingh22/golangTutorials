package main

import (
	"testing"
)

// Test names ALWAYS start with TestXXX(t *testing.T){}
func TestSum(t *testing.T) {
	// Method 1
	// t.Errorf("this test fails because i said so")
	// t.Fatalf("this test fails and stopped the execution proces.")

	// output := sum(1, 2, 3, 4, 5)
	// if output != 15 {
	// 	t.Errorf("Not working, as we got %v", output)
	// }

	// output = sum(1, -1)
	// if output != 0 {
	// 	t.Errorf("Not working, as we got %v", output)
	// }

	// Method 2: Table of Test, this reduces duplicate codes and makes testing elegant as tables.
	// tt := []struct {
	// 	nums []int
	// 	sum  int
	// }{
	// 	{[]int{1, 2, 3}, 6},
	// 	{[]int{1, -1}, 0},
	// 	{nil, 0},
	// }

	// for _, tc := range tt {
	// 	output := sum(tc.nums...)
	// 	if output != tc.sum {
	// 		t.Errorf("Not working, as we got %v", output)
	// 	}
	// }

	// Method 3: Name each sub-test for better documenting
	tt := []struct {
		name string
		nums []int
		sum  int
	}{
		{"Testing 3 digits", []int{1, 2, 3}, 6},
		{"Testing +1 & -1", []int{1, -1}, 0},
		{"passing nil", nil, 0},
	}

	for _, tc := range tt {
		// this creates a subtest which are all indepedent of each other
		t.Run(tc.name, func(t *testing.T) {
			output := Sum(tc.nums...)
			if output != tc.sum {
				t.Fatalf("Not working, as we got %v", output)
			}
		})
	}

}
