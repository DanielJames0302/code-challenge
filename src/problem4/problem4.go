package main

import "fmt"

// Implementation A: Iterative Sum (O(n))
func sum_to_n_a(n int) int {
    sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}

		return sum
}

// Implementation B: Mathematical Formula (O(1))
func sum_to_n_b(n int) int {
    return (n * (n + 1)) / 2
}

// Implementation C: Recursive Sum (O(n))
func sum_to_n_c(n int) int {
    if n == 0 {
        return 0
    }
    return n + sum_to_n_c(n-1)
}

func main() {
    fmt.Println(sum_to_n_a(5)) // Output: 15
    fmt.Println(sum_to_n_b(5)) // Output: 15
    fmt.Println(sum_to_n_c(5)) // Output: 15
}
