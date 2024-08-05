package testMap

import "fmt"

func TestMap() {
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap is null")
	}
	// 在使用map前需要先使用make给map 分配存储空间 动态分配存储空间 如果满了继续开辟该数量空间
	myMap = make(map[string]string, 10)

	myMap["one"] = "java"
	myMap["three"] = "c++"
	myMap["two"] = "python"
	fmt.Println(myMap)
	// myMap["four"] = "golang"

	// 方式2
	myMap1 := make(map[int]string)
	myMap1[1] = "java"
	myMap1[2] = "c++"
	myMap1[3] = "python"
	fmt.Println(myMap1)

	// 方式3
	myMap2 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "golang",
	}

	fmt.Println(myMap2)
	fmt.Println(myMap2["one"])

	// 遍历map
	for key, value := range myMap1 {
		fmt.Println("key: ", key, "value: ", value)
	}
	// 删除
	delete(myMap2, "three")
	fmt.Println("-----delete-----")
	printMap(myMap2)

	// 修改
	fmt.Println("-----update-----")
	pushValue(myMap2, "four", "java")
	printMap(myMap2)
}

// 参数调用
func printMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println("key: ", key, "value: ", value)
	}
}

// 修改map
func pushValue(myMap map[string]string, key string, value string) {
	myMap[key] = value
}
