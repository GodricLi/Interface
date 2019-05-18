package main

import (
	"fmt"
)

//定义一个接口类型
type Humaner interface {
	//方法，只有声明，没有实现，有别的类型（自定义类型）实现
	Sayhi()
}

type Students struct {
	Name string
}

// Students实现具体方法
func (stu *Students) Sayhi() {
	fmt.Println(stu.Name)
}

//
type School struct {
	Addr string
}

func (sc *School) Sayhi() {
	fmt.Println(sc.Addr)
}

type Mystr string
func (my *Mystr) Sayhi() {
	fmt.Println("i am a string")
}

//接口作为参数
func test_sayhi(i Humaner)  {
	i.Sayhi()
}




func main() {
	// 定义接口类型的变量
	var i Humaner
	// 只要实现了此接口方法的类型，那么这个类型的变量（此方法接收者类型）就可以给i赋值，
	// 这个类型接收的方法包含Humaner里面的方法
	// Students实现了接口的方法Sayhi，这个类型的变量stu就可以给i赋值，然后调用Sayhi，
	// 会执行Students的Sayhi的方法
	stu := &Students{
		Name: "Stu1",
	}
	i = stu
	i.Sayhi()

	sc := &School{"wuhan"}
	i = sc
	sc.Sayhi()

	var my Mystr = "str"
	i = &my
	i.Sayhi()

	//作为函数的参数，调用同一函数，不同的表现，多种状态，多态性
	test_sayhi(stu)
	test_sayhi(sc)
	test_sayhi(&my)

	//创建一个接口类型的切片
	x := make([]Humaner,3)
	x[0]=stu
	x[1]=sc
	x[2]=&my
	for _, v := range x{
		v.Sayhi()
	}
}
