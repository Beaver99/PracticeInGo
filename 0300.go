package lc300

import (
	"fmt"
	"math"
)

// ## challenging
// ### very challenging
//print the sub sequence ##
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	longest_len_sub := make([]int, len(nums))

	for i, _ := range nums {
		if i == 0 {
			longest_len_sub[i] = 1
			continue
		}
		// find j | j < i && nums[j] < nums[i] == m && longest_len_sub[j]
		max_sub_len, sub_tail := findFitSub(nums, longest_len_sub, i)
		if sub_tail != -1 {
			longest_len_sub[i] = max_sub_len + 1
		} else {
			longest_len_sub[i] = 1
		}
	}
	fmt.Println(longest_len_sub)
	return sliceMax(longest_len_sub)
}
func findFitSub(nums []int, longest_len_sub []int, index int) (max_sub_len int, sub_tail int) { //err
	max_sub_len = 1
	sub_tail = -1
	for j, _ := range nums {
		if j < index {
			if nums[j] < nums[index] && longest_len_sub[j] >= max_sub_len {
				max_sub_len = longest_len_sub[j]
				sub_tail = j
			}
		} else {
			break
		}
	}

	return max_sub_len, sub_tail
}

func sliceMax(nn []int) int {
	max := math.MinInt32
	for _, n := range nn {
		if n > max {
			max = n
		}
	}
	return max
}
func main() {
	a := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(a)
	fmt.Println(res)
}

func lengthOfLIS_others(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	longest_len_sub := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		longest_len_sub[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				longest_len_sub[i] = max(longest_len_sub[j]+1, longest_len_sub[i])
			}
		}
		result = max(result, longest_len_sub[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
