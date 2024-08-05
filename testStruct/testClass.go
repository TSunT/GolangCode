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
func (this *Hero) GetName() {
	fmt.Println("Name: ", this.Name)
}

// 首字母大写 表示公有 可以外部访问相当于java的public 否则只有包内部可访问
func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this *Hero) Show() {
	fmt.Println("{")
	fmt.Println("\tName:", this.Name)
	fmt.Println("\tAge:", this.Age)
	fmt.Println("\tLevel:", this.Level)
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
