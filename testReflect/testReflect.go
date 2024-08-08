package testreflect

import (
	"fmt"
	"reflect"
	"time"
)

func reflectField(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func reflectPrint(arg interface{}) {
	// 获取输入对象的type
	inputType := reflect.TypeOf(arg)
	fmt.Println("inputType is: ", inputType.Name())
	inputValue := reflect.ValueOf(arg)
	// 通过type 获取对象里面的属性 (访问公开属性)
	// 1. 获取interface的reflect.Type 通过type得到numField， 进行遍历
	// 2. 得到每一个field 获取其数据类型
	// 3. 通过value 获取field对应的value 使用 Interface()
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		if field.IsExported() { // 判定是否可导出 只要是 "public" 首字母大写的属性才可以判定是true
			fieldInfo := field.Tag.Get("info")
			value := inputValue.Field(i).Interface()
			fmt.Printf("field name: %s, type: %v, Tag: %s, value: %v \n", field.Name, field.Type, fieldInfo, value)
		} else {
			fmt.Println("field is Not Exported!")
		}

	}
	// 通过type 获取其方法 (访问公开方法)
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

type User struct {
	Id    string    `info:"Id" json:"id"`
	Age   int       `info:"Age" json:"age"`
	Name  string    `info:"Name" json:"name"`
	birth time.Time `info:"birth" json:"birth"`
}

func (user User) CallUser() {
	fmt.Println("call User...")
}

func TestReflectFunc() {
	var num float64 = 0.618

	reflectField(num)

	user := User{Id: "1", Age: 11, Name: "abc", birth: time.Now()}
	reflectPrint(user)
}
