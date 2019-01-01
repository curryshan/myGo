package main

import "fmt"

//  go中抛出一个panic异常 然后在defer中通过 recover捕获这个异常
func main() {

	defer func() {
        fmt.Println("c")
        //  分号表示&&
        if err := recover(); err != nil {
            fmt.Println(err)
        }
        fmt.Println("d")
	}()

	f()
}

func f () {
    fmt.Println("a")
    panic("this is error")
    fmt.Println("b")
    fmt.Println("f")
}

/*
a
c
this is error
d
*/
