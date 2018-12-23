package main

import (
    "fmt"
)

type PersonInfo struct {
    ID string
    Name string
    Address string
}

//  声明map
var personMap map[string] PersonInfo

//  声明并且初始化map
var personMoreMap = map[string] PersonInfo{
    "1234" : PersonInfo{"1234", "hao", "wuqiang..."},
}


func main() {

    //  使用make函数创建map 并在创建时制定map的存储能力
    personMap = make(map[string] PersonInfo, 100)

    //  向map中插入数据
    personMap["123"] = PersonInfo{"123", "shan", "raoyan..."}
    personMap["456"] = PersonInfo{"456", "meng", "anping..."}

    //  查找
    var person PersonInfo
    var ok bool


    person, ok = personMap["123123"]

    if ok {
        fmt.Println(person.Name)
    } else {
        fmt.Println("person not existed")
    }
    person, ok = personMap["123"]
    if ok {
        fmt.Println(person.Name)
    } else {
        fmt.Println("person not existed")
    }


    fmt.Println(personMap)

    //  元素删除
    delete(personMoreMap, "1234")
    fmt.Println(personMoreMap)


}

