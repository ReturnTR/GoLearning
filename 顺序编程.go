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

/*
v13 := 10    //同上，但是该句只能出现在函数体内，否则会报 syntax error: non-declaration statement outside function body，见main函数
*/

/*
//变量赋值，该句只能出现在函数体内，见main
v1="1"
*/

//常量,同样不允许重复赋值
const Pi float32 = 3.1415926
const c= 1  // 自动判断类型的方式，与变量不同，不用加冒号
const a,b=1,2 //多重赋值
//右边不能出现运行期才能知道的结果，const在编译时就已经确定好值了
//非法 const Pi := 3.1415926

//预定义常量，true，false，iota
/*
iota的使用：常量枚举
1.在每一个const关键字出现时被重置为0
2.每出现一次iota，其数字加1
*/
const (         //itoa被重置为0
    c0=iota     //c0 == 0
    c1=iota     //c1 == 1
    c2=iota     //c2 == 2
)
//如果两个const赋值语句是一样的，则可以省略后面的赋值表达式，
const (         //itoa被重置为0
    Sunday=iota     
    Monday          
    Tuesday
    Wednesday
    Thusday
    Friday
    Saturday
    number_of_days //开头小写为包内私有，其他符号可被其他包访问
)


//内置类型
/*
布尔：bool t只有两个值 rue false
数字：
    整数：int8 byte int16 int uint uintptr
    小数：float32 float64
    复数：complex64 complex128
字符：
单个字符：rune
字符串：string
错误类型：error
复合类型：
    指针：pointer
    数组：array
    切片：slice
    字典：map
    通道：chan
    结构体：struct
    接口：interface
    函数：func



*/


func main() {
    //变量初始化的简洁形式
    v13:=10
    name := "Hello World"
    // 变量声明后你必须对它做点什么，否则会报错:变量名 declared but not used，但不包括函数外的变量
    fmt.Println(name, v13)
    v13+=1   
    //变量赋值,不能赋值为其他类型的值，否则会报错：cannot use 赋值类型 (untyped string constant) as 定义类型 value in assignment
    v1=1 
    // 简洁的swap赋值(多重赋值)
    i:=0
    j:=1
    i,j=j,i
    fmt.Println(i,j)
    // 查看变量类型，同时还有reflect方法
    fmt.Printf("%T\n",c1)
    var cc rune = 'c'
    fmt.Printf("%T\n",cc)  // rune实则为int32类型

    //强制类型转换
    v12=int(v11)
    


}