package main

import (
	"fmt"
)

func main() {
	prices := []int{1, 5, 8, 9, 10}
	n := 4
	fmt.Printf("Maximum Obtainable Value is %d\n", cutRod(prices, n))
	fmt.Printf("Maximum Obtainable Value is %d\n", memoizedCutRod(prices, n))
	fmt.Printf("Maximum Obtainable Value is %d\n", bottomUpCutRod(prices, n))
}

func cutRod(p []int, n int) int {
	if n == 0 {
		return 0
	}

	q := -1

	for i := 1; i <= n; i++ {
		if i < len(p) {
			q = max(q, p[i-1]+cutRod(p, n-i))
		}
	}

	return q
}

/*
Memoization
*/
func memoizedCutRod(p []int, n int) int {
	r := make([]int, n+1)

	for i := 0; i < n; i++ {
		r[i] = -1
	}

	return memoizedCutRodAux(p, n, r)
}

func memoizedCutRodAux(p []int, n int, r []int) int {
	if r[n] > 0 {
		return r[n]
	}

	var q int

	if n == 0 {
		q = 0
	} else {
		q = -1
		for i := 1; i <= n; i++ {
			q = max(q, p[i-1]+memoizedCutRodAux(p, n-i, r))
		}
	}

	r[n] = q

	return q
}

/*
Bottom-up method
*/
func bottomUpCutRod(p []int, n int) int {
	r := make([]int, n+1)

	r[0] = 0

	for j := 1; j <= n; j++ {
		q := -1
		for i := 1; i <= j; i++ {
			q = max(q, p[i-1]+r[j-i])
		}
		r[j] = q
	}

	return r[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
