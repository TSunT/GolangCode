package leetCode

import "fmt"

func lengthOfLIS(nums []int) int {
	fmt.Println(nums)
	length := len(nums)
	if length < 2 {
		return length
	}

	maxLen := 1
	// 构建dp二维数组
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	// 初始化 默认单一元素一定是递增的
	for i := 0; i < length; i++ {
		dp[i][i] = true
	}
	for i := 0; i < length; i++ {
		// 以i列从小到i进行遍历
		for j := 0; j < i; j++ {
			// 需要判定的最大下标
			cursor := i + j - 1
			// fmt.Println(i, "列", j, "行", cursor, "最大下标")
			if cursor >= length {
				break
			}
			if nums[j] >= nums[cursor] {
				dp[j][cursor] = false
			} else {
				if cursor-j < 3 {
					dp[j][cursor] = true
				} else {
					dp[j][cursor] = dp[j+1][cursor-1] && nums[j] < nums[j+1] && nums[cursor-1] < nums[cursor]
				}
			}
			if dp[j][cursor] && cursor-j+1 > maxLen {
				maxLen = cursor - j
			}
		}
	}
	return maxLen
}
