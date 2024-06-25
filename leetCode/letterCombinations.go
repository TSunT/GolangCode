package leetCode

import "fmt"

func letterCombinations(digits string) []string {

	// 构建字符映射关系
	var digitCharMap map[string]string = make(map[string]string)
	digitCharMap["2"] = "abc"
	digitCharMap["3"] = "def"
	digitCharMap["4"] = "ghi"
	digitCharMap["5"] = "jkl"
	digitCharMap["6"] = "mno"
	digitCharMap["7"] = "pqrs"
	digitCharMap["8"] = "tuv"
	digitCharMap["9"] = "wxyz"

	res := make([]string, 0)
	length := len(digits)
	for i := 0; i < length; i++ {
		var digitNum string = digits[i : i+1]
		fmt.Println(digitNum)
		btnStr := digitCharMap[digitNum]
		res = append(res, btnStr)
	}
	return res
}
