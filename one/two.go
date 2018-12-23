package main

import "fmt"

var str1 string
var str2 string


func main() {
    str1 = "shanmenghao"
    str2 = "shuaibi"
    firstChar := str1[0]

    fmt.Printf("string len is : %d\n", len(str1))
    fmt.Printf("string first char is : %c\n", firstChar)
    fmt.Printf("str1 + str2 is : %s\n", str1 + str2)

    //  字符串遍历 方法一
    //length := len(str1)
    //for i := 0; i < length; i++ {
    //    fmt.Printf("%c\n", str1[i])
    //}

    //  字符串遍历 方法二
    for i, _ := range str1 {
        fmt.Printf("%c\n",str1[i])
    }


}

