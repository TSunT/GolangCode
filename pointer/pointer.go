package main

import "fmt"

func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Println("a = ", a, "b = ", b)
	swapPointer(&a, &b)
	fmt.Println("a = ", a, "b = ", b)
	var pa *int = &a

	fmt.Println(pa)
	fmt.Println(&a)

	var secondLevelP **int = &pa

	fmt.Println(&pa)
	fmt.Println(secondLevelP)

}

func swap(a int, b int) {
	c := a
	a = b
	b = c
}

func swapPointer(pa *int, pb *int) {
	var temp int = *pa
	*pa = *pb
	*pb = temp
}
