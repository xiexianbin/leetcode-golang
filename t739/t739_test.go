package t739

import (
	"fmt"
	"testing"
)

func dailyTemperatures(T []int) []int {
	result := []int{}
	for i, v := range T {
		stack := []int{}
		j := i
		for ; j < len(T); j++ {
			if v < T[j] {
				result = append(result, len(stack))
				break
			} else {
				stack = append(stack, v)
			}
		}
		if j == len(T) {
			result = append(result, 0)
		}
	}
	return result[:len(T)]
}

/* https://leetcode-cn.com/explore/learn/card/queue-stack/218/stack-last-in-first-out-data-structure/879/ */
func TestT879(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
