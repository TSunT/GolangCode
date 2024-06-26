package leetCode

/*
	17.电话号码的字母组合
	标准回溯法
*/

var digitCharList []string = []string{
	"",    // 0
	"",    // 1
	"abc", // 2
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinationsBackTrack(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	backTrack(&digits, 0, &res, "")
	return res
}

// 回溯法
func backTrack(digitsPtr *string, cursor int, result *[]string, ancestorStr string) {
	// 递归到最底层
	if cursor >= len(*digitsPtr) {
		*result = append(*result, ancestorStr)
		return
	}
	// 不是最底层 获取输入的该字符 并获得该字符对应的字符串
	btnOffset := int((*digitsPtr)[cursor]) - int('0')
	currentLevelStr := digitCharList[btnOffset]
	for i := 0; i < len(currentLevelStr); i++ {
		// 遍历每个字符 递归拼接
		backTrack(digitsPtr, cursor+1, result, ancestorStr+string(currentLevelStr[i]))
	}
}
