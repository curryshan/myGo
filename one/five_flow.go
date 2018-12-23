package main

import "fmt"


func main() {

	retOne := testOne(4)
	fmt.Println(retOne)

	retTwo := testTwo(5)
	fmt.Println(retTwo)

	retThree := testThree()
	fmt.Println(retThree)

	retFour := testFour()
	fmt.Println(retFour)

	testFive()

	retSix := testSix()
	fmt.Println(retSix)

	v1 := 1000
	v2 := "shan"
	v3 := false

	testSeven(v1, v2, v3)
}


func testOne (x int) bool {
	//  条件控制语句
	if x > 5 {
		return true
	} else {
		return false
	}
}

func testTwo (x int) string {
	//  不需要break关键字退出 默认退出 如不退出 fallthrough
	var ret string

	switch x {
		case 0 :
			ret = "0"
		case 1:
			fallthrough
		case 2:
			ret = "2"
		case 3:
			ret = "3"
		default:
			ret = ">3"
	}

	return ret
}

func testThree () int {
	sum := 0
	for i:=0; i<10; i++ {
		sum += 1
	}
	return sum
}

func testFour() int {
	//  无限循环场景
	sum := 0
	for {
		sum ++
		if sum > 100 {
			break
		}
	}
	return sum
}

func testFive() {
	//  break高级用法

	JLoop:
	for i:=0; i<20; i++ {
		for j:=0; j<20; j++ {
			if j>5 {
				//  这里跳出了所有循环
				break JLoop
			}
			fmt.Printf("i = %d\tj = %d\n", i, j)
		}
	}
}

func testSix() int {
	i := 0
	HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
	return i
}

func testSeven(args ...interface{}) {
	//  不定参数
	for _, val := range args {
		fmt.Println(val)
	}
}