package twopointers

import "fmt"

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		left, right := 0, n-1
		for left < right {
			if knows(left, right) {
				left++
			} else {
				right--
			}
		}
		fmt.Println(left, right)
		for i := 0; i < n; i++ {
			if i != right && (knows(right, i) || !knows(i, right)) {
				return -1
			}
		}
		return right
	}
}
