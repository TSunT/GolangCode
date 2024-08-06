package testStruct

import "fmt"

// 定义一个对象
type Hero struct {
	// 首字母大写 表示公有 可以外部访问相当于java的public 否则只有包内部可访问
	Name  string
	Age   int
	Level int
}

// this Hero  是对该对象的拷贝 只有 this *Hero才是对于原有的对象的操作
func (hero *Hero) GetName() {
	fmt.Println("Name: ", hero.Name)
}

// 首字母大写 表示公有 可以外部访问相当于java的public 否则只有包内部可访问
func (hero *Hero) SetName(name string) {
	hero.Name = name
}

func (hero *Hero) Show() {
	fmt.Println("{")
	fmt.Println("\tName:", hero.Name)
	fmt.Println("\tAge:", hero.Age)
	fmt.Println("\tLevel:", hero.Level)
	fmt.Println("}")
}

func TestHeroClass() {
	hero := Hero{Name: "zhang3", Age: 12, Level: 13}
	// 查看
	hero.Show()
	fmt.Println("----changeName----")
	hero.SetName("li4")
	hero.GetName()
}
