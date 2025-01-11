package main

func Add(a int, b int) int {
	return a + b
}

func Factorial(n int) (int, error) {
	if n == 0 {
		return 1, nil
	}
	if n < 0 {
		return 0, nil
	}

	result, _ := Factorial(n - 1)
	return n * result, nil
}

// func fibonacci(n int) (int, int) {
// 	if n <= 1 {
// 		return n, n
// 	}

// 	_, prev1 := fibonacci(n - 1)
// 	_, prev2 := fibonacci(n - 2)

// 	return n, prev1 + prev2
// }

// func fibonacci2(n int) int {
// 	c := 0
// 	prev1, prev2, result := 1, 1, 0
// 	for c < n {
// 		prev1 = prev2
// 		prev2 = result
// 		result = prev1 + prev2
// 		c += 1
// 	}
// 	return result
// }
