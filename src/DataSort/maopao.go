/*
冒泡排序是效率较低的排序算法，可以说是最慢的排序算法了，我们只需知道它是什么，在实际工作上切勿使用如此之慢的排序算法!

*/

package main

import "fmt"

func BubbleSort(list []int) {
    n := len(list)
    // 在一轮中有没有交换过
    didSwap := false

    // 进行 N-1 轮迭代
    for i := n - 1; i > 0; i-- {
        // 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
        for j := 0; j < i; j++ {
            // 如果前面的数比后面的大，那么交换
            if list[j] > list[j+1] {
                list[j], list[j+1] = list[j+1], list[j]
                didSwap = true
            }
        }

        // 如果在一轮中没有交换过，那么已经排好序了，直接返回
        if !didSwap {
            return
        }
    }
}

func main() {
    list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    BubbleSort(list)
    fmt.Println(list)
}