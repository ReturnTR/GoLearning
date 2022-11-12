package main

import "fmt"

//go与C一样，是静态类型语言，需要先编译整个文件，然后再运行

// 运行程序： go run main.go

// 变量声明
var v1 int
/*
//该句会报错，不允许重复声明
var v1 string 
*/
var v2 string
var v3 [10]int
var v4 []int        //数组切片
var v5 struct{
    f int 
}
var v6 *int  //指针
var v7 map[string]int   //map key为string类型，value为int类型，类似字典
var v8 func(a int) int  //函数类型，前一个int为参数类型，后一个int为返回类型

var (       // 避免var重复
    v9 int
    v10 string
)

//变量初始化，go语言有闭包原则，不能在函数体外赋值
var v11 int = 10
var v12 = 10  //编译器自动推导出v2类型
v13 := 10    //同上，但是该句只能出现在函数体内，否则会报 syntax error: non-declaration statement outside function body

// //已经声明过的变量不能再重复声明

// //变量赋值
// var v1 string
// v1="1"


func main() {
    // name := "Go Developers"
    fmt.Println("Azure for", v11)
}