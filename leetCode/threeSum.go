package leetCode

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// fmt.Println(nums)
	res := make([][]int, 0)
	length := len(nums)
	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		end := length - 1
		for j := i + 1; j < length-1; j++ {
			// 判定是否重复
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			a := nums[i]
			b := nums[j]
			// if a+b+nums[j+1] > 0 {
			// 	break
			// }
			for end > j && a+b+nums[end] > 0 {
				end--
			}
			c := nums[end]
			// if end < length-1 && nums[end] == nums[end + 1] {
			// 	continue
			// }
			if end <= j {
				break
			}

			// fmt.Println("i: ", i, "j: ", j, "end: ", end, "-> a: ", a, "b: ", b, "c: ", c)
			if (a + b + c) == 0 {
				var intArrList []int = []int{a, b, c}
				res = append(res, intArrList)
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	res := make([][]int, 0)
	length := len(nums)
	for i := 0; i <= length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j <= length-2; j++ {
			end := length - 1
			third := j + 1
			// 判定是否重复
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			a := nums[i]
			b := nums[j]

			for end > third {
				if end < length-1 && nums[end] == nums[end+1] {
					end--
					continue
				}
				if third > j+1 && nums[third] == nums[third-1] {
					third++
					continue
				}
				if a+b+nums[third]+nums[end] > target {
					end--
				}
				if a+b+nums[third]+nums[end] < target {
					third++
				}
				if end <= third {
					break
				}
				if a+b+nums[third]+nums[end] == target {

					fmt.Println(" i: ", i, " j: ", j, " third: ", third, " end: ", end, "-> a: ", a, " b: ", b, " c: ", nums[third], " d: ", nums[end])
					var intArrList []int = []int{a, b, nums[third], nums[end]}
					res = append(res, intArrList)
					end--
					third++
				}
			}

			if end <= third {
				continue
			}

			// c := nums[third]
			// d := nums[end]
			// if (a + b + c + d) == target {
			// 	var intArrList []int = []int{a, b, c, d}
			// 	res = append(res, intArrList)
			// }
		}
	}
	return res
}
