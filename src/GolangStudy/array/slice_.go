package main

import "fmt"


//
func a(mySlice []int) {
	mySlice[0] = 2
	fmt.Println(mySlice)
}

func main() {
	//切片有常见的两种声明方式，有初始值和没有初始值
	// myArray := []int{0,0}
	myArray := make([]int, 3)
	//动态数组容量增加和截取
	//如果你不指定cap(容量)，默认容量的基数就是长度，容量是以长度基数为基准，成倍增加
	for i := 0; i < 19; i ++ {
		myArray = append(myArray, 1)
	}
	a(myArray)
	//go的切片基本和python切片一摸一样,但是注意，截取出来的切片内部变化会影响母切片，这跟python是不一样的
	a1 := myArray[:3]
	a2 := make([]int, 3)
	
	//将后者的值按顺序拷贝到前者的值当中
	copy(myArray, a2)
	fmt.Println(myArray, len(myArray),cap(myArray), a1, a2)
}

// 平时我们很少用数组，基本都是用动态数组(切片)
// 因为如果用数组，在函数传参的时候必须严格传规定长度的数组，且数组是值拷贝
// 再强调一遍，数组来讲[10]string 和 [5]string是不同类型

