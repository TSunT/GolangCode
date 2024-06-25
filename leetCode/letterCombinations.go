package leetCode

import "fmt"

/*
   17.电话号码的字母组合
   给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
   给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
   示例 1：

	输入：digits = "23"
	输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
	示例 2：

	输入：digits = ""
	输出：[]
	示例 3：

	输入：digits = "2"
	输出：["a","b","c"]
*/
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
	// 遍历输入的字符
	for i := 0; i < length; i++ {
		var digitNum string = digits[i : i+1]
		fmt.Println(digitNum)
		btnStr := digitCharMap[digitNum]
		if btnStr != "" {
			charList := make([]string, 0)
			if len(res) > 0 {
				// 存在先前的字符
				for j := 0; j < len(res); j++ {
					for k := 0; k < len(btnStr); k++ {
						charList = append(charList, res[j]+btnStr[k:k+1])
					}
				}
				res = charList
			} else {
				// 不存在先前的字符
				for j := 0; j < len(btnStr); j++ {
					charList = append(charList, btnStr[j:j+1])
				}
				res = charList
			}
		}
	}
	return res
}
