package main

import (
	// leetCodeLib "learnCode/helloworld/leetCode"
	// mylib1 "learnCode/helloworld/lib1" // mylib1 给包取个别名 import 进来后会首先调用该模块下的init方法
	// mylib2 "learnCode/helloworld/lib2"
	// testMap "learnCode/helloworld/testMap"
	// testStruct "learnCode/helloworld/testStruct"
	// testFunction "learnCode/helloworld/function"
	// testInterface "./testInterface"
	// testPair "./testPair"
	// testReflect "./testReflect"
	// testGoroutine "./testGoroutine"
	testChannel "./testChannel"
)

func main() {
	// mylib1.Lib1_test()
	// mylib2.Lib2_test()
	// leetCodeLib.Test()
	// testMap.TestMap()
	// testStruct.TestHeroClass()
	// testStruct.TestHumanInherit()
	// testFunction.TestFunc()
	// testInterface.TestInterfaceAssert()
	// testReflect.TestStructToJson()
	// testReflect.TestJsonToStruct()
	// testGoroutine.TaskMainTask()
	testChannel.Text_Fibonacii()
	testChannel.TestChannelWithCache()
}
