package main

import "fmt"

func main() {
	var slice0 []int
	if slice0 == nil {
		fmt.Println("slice0 是个空切片，未开辟内存空间")
		fmt.Println("slice0 len = ", len(slice0), "cap = ", cap(slice0), "value = ", slice0)
	}
	fmt.Println("-------------------")
	slice1 := make([]int, 3)
	fmt.Println("slice1 len = ", len(slice1), "cap = ", cap(slice1), "value = ", slice1)
	fmt.Println("-------------------")
	slice2 := []int{1, 2, 3}
	fmt.Println("slice2 len = ", len(slice2), "cap = ", cap(slice2), "value = ", slice2)
	fmt.Println("--------元素追加 如果 当前cap 大于于 len cap空间不做变化 同时len加1-----------")
	slice2 = append(slice2, 4)
	fmt.Println("slice2 len = ", len(slice2), "cap = ", cap(slice2), "value = ", slice2)
	fmt.Println("---------当追加时 如果cap 不够了则go语言自动开辟当前cap元素个数的存储空间 供append 使用 ----------")
	slice2 = append(slice2, 5)
	fmt.Println("slice2 len = ", len(slice2), "cap = ", cap(slice2), "value = ", slice2)
	fmt.Println("-------------------")
	slice2 = append(slice2, 6)
	fmt.Println("slice2 len = ", len(slice2), "cap = ", cap(slice2), "value = ", slice2)
	fmt.Println("-------------------")
	slice2 = append(slice2, 7)
	fmt.Println("slice2 len = ", len(slice2), "cap = ", cap(slice2), "value = ", slice2)
	fmt.Println("--------元素截取 实际上相当于SQL的视图 不是copy-----------")

	sub_slice2 := slice2[0:3]
	fmt.Println("sub_slice2 len = ", len(sub_slice2), "cap = ", cap(sub_slice2), "value = ", sub_slice2) // 切片截取 [0:3]  表示截取第0~2（不包括3）号下标的元素
	sub_slice2_1 := slice2[3:]
	fmt.Println("sub_slice2_1 len = ", len(sub_slice2_1), "cap = ", cap(sub_slice2_1), "value = ", sub_slice2_1) // 切片截取 [3:]  表示截取第3号下标及其后续的全部的元素
	fmt.Println("------ 修改截取切片的数值 改变原来切片的数值------------")
	sub_slice2[0] = 111 // 截取切边改变值实际上改变了原有切片的值
	fmt.Println("sub_slice2 len = ", len(sub_slice2), "cap = ", cap(sub_slice2), "value = ", sub_slice2)
	fmt.Println("slice1 len = ", len(slice2), "cap = ", cap(slice2), "value = ", slice2)
}
