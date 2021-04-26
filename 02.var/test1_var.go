package main

import "fmt"

/*
	四种变量的声明方式
*/

//声明全局变量前三种是可以的,第四种不支持。省去var的写法只能在函数体中使用
var ga int

func main() {
	//方法一：声明一个变量,默认的值是零
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)
	var aa float32
	fmt.Println("aa = ", aa)
	fmt.Printf("type of aa = %T\n", aa)
	var aaa string
	fmt.Println("aaa = ", aaa)
	fmt.Printf("type of aaa = %T\n%s\n", aaa, aaa)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc string = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四：省去var关键字（最常的方法）
	d := "ccdd"
	fmt.Println("d =", d)
	fmt.Printf("type of d = %T\n", d)

	//声明多个变量
	//此处是同为int提取出来，而不是只为yy声明
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, ", yy =", yy)

	var kk, ll = 100, "aceld"
	fmt.Println("kk =", kk, ", ll =", ll)
	fmt.Printf("type of kk = %T, ll = %T\n", kk, ll)

	var (
		pp = "x1"
		oo = "x2"
	)
	fmt.Println("pp =", pp, ", oo =", oo)
	fmt.Printf("type pp %T, type oo %T\n", pp, oo)

	ii, uu := "x3", 10
	fmt.Println("ii =", ii, ", uu =", uu)
	fmt.Printf("type ii %T, type uu %T\n", ii, uu)
}
