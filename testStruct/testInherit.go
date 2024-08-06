package testStruct

import "fmt"

type Human struct {
	name   string
	gender string
	age    int8
}

type Superman struct {
	Human // 表示Superman 继承了Human里面的属性和方法
	level int
}

func (human *Human) Eat() {
	fmt.Println("Human: ", human.name, ",Eat...")
}

func (human *Human) Walk() {
	fmt.Println("Human: ", human.name, ",Walk...")
}

// 子类中重写父类的方法
func (superman *Superman) Eat() {
	fmt.Println("superman: ", superman.name, ",level:", superman.level, ",Eat...")
}

// 父类中定义自己的方法
func (superman *Superman) Fly() {
	fmt.Println("superman: ", superman.name, ",level:", superman.level, ",Fly...")
}

// 子类方法中调用父类属性和方法
func (superman *Superman) Print() {
	superman.Walk()
	fmt.Println("{")
	fmt.Println("\tname:", superman.name)
	fmt.Println("\tgender:", superman.gender)
	fmt.Println("\tlevel:", superman.level)
	fmt.Println("\tage:", superman.age)
	fmt.Println("}")
}

func TestHumanInherit() {
	h := Human{name: "zhang3", gender: "male", age: 23}
	h.Eat()
	h.Walk()
	fmt.Println("----superman----")

	// 定义一个子类 方式一 {}内包含子类的定义
	s := Superman{Human: Human{name: "li4", gender: "female", age: 22}, level: 10}
	s.Eat()
	s.Fly()
	// 定义一个子类 方式二
	var s2 Superman
	s2.name = "wang5"
	s2.gender = "male"
	s2.age = 27
	s2.level = 20

	s2.Print()
}
