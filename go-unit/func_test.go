package main

import (
	"testing"
)

func TestAdd(t *testing.T){
	// Mutiple testcases
	testcasesAdd := []struct{
		name     string
		a,b      int
		expected int
	}{
		{"Add positive numbers", 2, 3, 5},
		{"Add negative numbers", -2, -3, -5},
		{"Add zeros", 0, 0, 0},
	}

	for _, tc := range testcasesAdd{
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d is incorrect, an expected result is %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestFac(t *testing.T){
	// Mutiple testcases
	testcasesFact := []struct{
		name     string
		num      int
		expected int
	}{
		{"Case #1 -> 2!", 2, 2},
		{"Case #2 -> 5!", 5, 120},
	}

	for _, tc := range testcasesFact{
		t.Run(tc.name, func(t *testing.T) {
			result,_ := Factorial(tc.num)
			if result != tc.expected {
				t.Errorf("Factorial(%d) = %d is incorrect, an expected result is %d", tc.num, result, tc.expected)
			}
		})
	}
}