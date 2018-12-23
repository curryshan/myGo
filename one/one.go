package main

import "fmt"

//  声明变量
var (
    v1 int
    v2 string
)
var v3 [10]int
var v4 []int
var v5 struct {
    f int
}
var v6 *int
var v7 map[string]int //map类型 key为string value为int
var v8 func(a int) int

//  声明变量时需初始化
var v9 int = 10
var v10 = 10

//  常量定义
const Pi float64 = 3.1415926
const zero = 0.0    //无类型声明
const const_a, const_b, const_c = 3, 4, "string"//无类型 多重声明

//  预定义常量 iota 常量计数器
const (
    c1 = iota
    _
    _
    c3
)


//  不同类型的整形变量
var v21 int8 = 123
var v22 int = 456

const const_1  = 123;

func main() {

    v11 := 10//只可以用在函数体内

    //  函数多返回值 且可忽略
    _, _, nickName := getName()

    fmt.Println(nickName)
    fmt.Println(v11)

    //  常量计数器
    fmt.Println(c1)
    fmt.Println("result:", c3)

    //  常量对比变量
    sign := (v21 == const_1)
    fmt.Println(sign)


}

func getName() (firstName, lastName, nickName string){
    return  "shan", "meng", "hao"
}
