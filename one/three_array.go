package main

import "fmt"


//  数组 go中数组长度在定义后就不可更改
//  数组是一个值类型 参数传递时会发生复制
var arr1 [10] int

//  数组切片
//  数组切片的数据结构可以抽象为以下三个变量
//  一个指向原生数组的指针
//  数组切片中的元素个数
//  数组切片已分配的存储空间
//  可动态增减元素是数组切片比数组更强大的功能



func main() {

    arr1[0] = 10
    arr1[2] = 8
    arr1[4] = 9

    for i, v := range arr1 {
        fmt.Printf("arr1 index %d is %d\n", i, v)
    }

    //  验证数组的值传递特性
    arr2 := [5]int{1,2,3,4,5}
    modify(arr2)
    fmt.Println(arr2)


    //  创建数组切片 方法一 基于数组创建
    var arr1Slice1 []int = arr1[:]
    var arr1Slice2 []int = arr1[1:4]
    var arr1Slice3 []int = arr1[:3]

    fmt.Println(arr1Slice1, arr1Slice2, arr1Slice3)

    //  直接创建数组切片
    arr1Slice4 := make([]int, 5)//创建一个初始元素个数为5的数组切片 元素初始值为0
    arr1Slice5 := make([]int, 5, 20)//创建一个初始元素个数为5的数组切片 元素初始值为0 并预留10个元素的存储空间
    arr1Slice6 := []int {1,2,3,4,5}//直接创建并初始化包含5个元素的数组切片

    arr1Slice7 := append(arr1Slice6, 6, 7, 8)
    arr1Slice8 := append(arr1Slice4, arr1Slice5...)

    //  cap函数返回的是 数组切片分配的空间大小
    fmt.Println(cap(arr1Slice4), cap(arr1Slice5), cap(arr1Slice6), cap(arr1Slice7), cap(arr1Slice8))

    //  数组切片复制 copy
    copy(arr1Slice1, arr1Slice2)//只会复制arr1Slice2的三个元素到arr1Slice1的前三个位置
    fmt.Println(arr1Slice1, arr1Slice2)

    copy(arr1Slice2, arr1Slice1)//只会复制arr1Slice1的前三个元素到arr1Slice2中
    fmt.Println(arr1Slice1, arr1Slice2)



}

func modify(arr [5]int) {
    arr[0] = 10
    fmt.Println(arr)
}

