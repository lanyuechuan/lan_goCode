/*
按照之前分析的几种排序算法，一般建议待排序数组为小规模情况下使用直接插入排序，
在规模中等的情况下可以使用希尔排序，但在大规模还是要使用快速排序，归并排序或堆排序。

*/

package main

import "fmt"

// 增量序列折半的希尔排序
func ShellSort(list []int) {
    // 数组长度
    n := len(list)

    // 每次减半，直到步长为 1
    for step := n / 2; step >= 1; step /= 2 {
        // 开始插入排序，每一轮的步长为 step
        for i := step; i < n; i += step {
            for j := i - step; j >= 0; j -= step {
                // 满足插入那么交换元素
                if list[j+step] < list[j] {
                    list[j], list[j+step] = list[j+step], list[j]
                    continue
                }
                break
            }
        }
    }
}

func main() {
    list := []int{5}
    ShellSort(list)
    fmt.Println(list)

    list1 := []int{5, 9}
    ShellSort(list1)
    fmt.Println(list1)

    list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    ShellSort(list2)
    fmt.Println(list2)

    list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}
    ShellSort(list3)
    fmt.Println(list3)
}