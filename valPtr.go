// 测试结构体初始化方式，使用传值和传指针方法区别。
// 下面例子同样可以将方法的接收者变更为指针类型再做测试。

package main

import "fmt"

type person struct {
	Name string
}

func (per person) EchoName() string {
	return string(per.Name)
}

func main() {

	// 先定义，再初始化，然后打印
	var PerBob person
	PerBob.Name = "Bob"
	fmt.Printf("%v\n", PerBob)

	// 通过传递值方式调用EchoName（）方法
	ValEcho := PerBob.EchoName()
	fmt.Printf("%v\n", ValEcho)

	var PerMike = &person{"Mike"}
	// 直接打印初始化的PerMike结构,输出应为&{Mike}
	fmt.Printf("%v\n", PerMike)

	// 通过传递指针方式调用EchoName（）方法
	PtrEcho := PerMike.EchoName()
	fmt.Printf("%v\n", PtrEcho)

}
