package testInterface

import "fmt"

// 接口已指针的形式使用 需要传递该对象的地址
type AnimalIf interface {
	Sleep()
	GetColor() string
	GetType() string
	GetAge() int
}

// cat
type Cat struct {
	color string
	age   int
}

func (cat *Cat) Sleep() {
	fmt.Println("cat Sleep...")
}

func (cat *Cat) GetColor() string {
	// fmt.Println("cat color is:", cat.color)
	return cat.color
}

func (cat *Cat) GetType() string {
	// fmt.Println("cat")
	return "cat"
}

func (cat *Cat) GetAge() int {
	return cat.age
}

// dog
type Dog struct {
	color string
	age   int
}

func (dog *Dog) Sleep() {
	fmt.Println("dog Sleep...")
}

func (dog *Dog) GetColor() string {
	// fmt.Println("dog color is:", dog.color)
	return dog.color
}

func (dog *Dog) GetType() string {
	// fmt.Println("dog")
	return "dog"
}

func (dog *Dog) GetAge() int {
	return dog.age
}

// 使用interface 实现多态调用
func PrintAnimal(animal AnimalIf) {
	fmt.Println("{")
	fmt.Println("\ttype:", animal.GetType())
	fmt.Println("\tcolor:", animal.GetColor())
	fmt.Println("\tage:", animal.GetAge())
	fmt.Println("}")
}
func TestInterface() {
	var animal AnimalIf
	animal = &Cat{color: "black", age: 2} // 注意 如果使用interface 则必须要全部实现接口方法 且返回类型一致，否则不合法
	animal.Sleep()
	fmt.Println("animal type: ", animal.GetType(), "animal age: ", animal.GetAge())
	animal = &Dog{color: "brown", age: 3}
	animal.Sleep()
	fmt.Println("animal type: ", animal.GetType(), "animal age: ", animal.GetAge())
	fmt.Println("----Polymorphic----")
	PrintAnimal(&Cat{color: "gray", age: 1})
	PrintAnimal(&Dog{"white", 4})
}
