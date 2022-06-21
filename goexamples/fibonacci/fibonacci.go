package fibonacci

import (
	"context"
	"fmt"
	"time"
)

// Create a function that returns list of N consecutive
// fibbonaci sequence numbers.
func fibbonaciNums(count int) []int {
	if count == 1 {
		return []int{0}
	}
	result := make([]int, count)
	a, b := 0, 1
	for i := 0; i < count; i++ {
		result[i] = a
		a, b = b, a+b
	}
	return result
}

// Create a function that returns list of N consecutive
// fibbonaci sequence numbers lower then the given limit.
func fibbonaciLimit(limit int) []int {
	if limit == 1 {
		return []int{0, 1, 1}
	}
	var result []int
	a, b := 0, 1
	for a < limit {
		result = append(result, a)
		a, b = b, a+b
	}
	return result
}

// Create a function that returns list of N consecutive
// fibbonaci sequence numbers lower then the given limit and times out after given Time.
func fibbonaciLimitTimeout(limit int, timeout time.Duration) []int {
	if limit == 1 {
		return []int{0, 1, 1}
	}
	var result []int
	ch := make(chan []int)
	go func(ch chan []int) {
		a, b := 0, 1
		for a < limit {
			result = append(result, a)
			a, b = b, a+b
		}
		ch <- result
	}(ch)
	select {
	case res := <-ch:
		return res
	case <-time.After(timeout):
		fmt.Print("Timed out: ")
		return []int{}
	}
}

// Create a function that returns list of N consecutive
// fibbonaci sequence numbers lower then the given limit and times out after given Time (using context).
func fibbonaciLimitTimeoutContext(ctx context.Context, limit int) []int {
	if limit == 1 {
		return []int{0, 1, 1}
	}
	var result []int
	ch := make(chan []int)
	go func(ch chan []int) {
		a, b := 0, 1
		for a < limit {
			result = append(result, a)
			a, b = b, a+b
		}
		ch <- result
	}(ch)
	select {
	case res := <-ch:
		return res
	case <-ctx.Done():
		fmt.Print("Canceled or timed out")
		return []int{}
	}
}

// Recursive fibbonaci, return nth fibonacci number
func fibbonaciRecursive(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibbonaciRecursive(num-1) + fibbonaciRecursive(num-2)
	}
}
