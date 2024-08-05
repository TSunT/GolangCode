package main

import (
	// leetCodeLib "learnCode/helloworld/leetCode"
	// mylib1 "learnCode/helloworld/lib1" // mylib1 给包取个别名 import 进来后会首先调用该模块下的init方法
	// mylib2 "learnCode/helloworld/lib2"
	testMap "learnCode/helloworld/testMap"
)

func main() {
	// mylib1.Lib1_test()
	// mylib2.Lib2_test()
	// leetCodeLib.Test()
	testMap.TestMap()
}
