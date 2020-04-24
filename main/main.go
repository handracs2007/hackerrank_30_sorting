// https://www.hackerrank.com/challenges/30-sorting/problem

package main

import (
	"fmt"
)

func bubbleSort(elements []int) (numSwaps, firstElement, lastElement int) {
	for i := 0; i < len(elements); i++ {
		for j := len(elements) - 1; j > 0; j-- {
			if elements[j] < elements[j-1] {
				// Swap
				var tmp = elements[j]
				elements[j] = elements[j-1]
				elements[j-1] = tmp

				numSwaps++
			}
		}
	}

	return numSwaps, elements[0], elements[len(elements)-1]
}

func main() {
	var n int
	_, _ = fmt.Scanln(&n)

	var elements = make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&elements[i])
	}

	var numOfSwaps, firstElement, lastElement = bubbleSort(elements)

	fmt.Printf("Array is sorted in %d swaps.\n", numOfSwaps)
	fmt.Printf("First Element: %d\n", firstElement)
	fmt.Printf("Last Element: %d\n", lastElement)
}
