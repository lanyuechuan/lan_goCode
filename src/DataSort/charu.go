/*
数组规模 n 较小的大多数情况下，我们可以使用插入排序，它比冒泡排序，选择排序都快，甚至比任何的排序算法都快。

数列中的有序性越高，插入排序的性能越高，因为待排序数组有序性越高，插入排序比较的次数越少。

大家都很少使用冒泡、直接选择，直接插入排序算法，因为在有大量元素的无序数列下，这些算法的效率都很低。

*/

package main

import "fmt"

func InsertSort(list []int) {
    n := len(list)
    // 进行 N-1 轮迭代
    for i := 1; i <= n-1; i++ {
        deal := list[i] // 待排序的数
        j := i - 1      // 待排序的数左边的第一个数的位置

        // 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
        if deal < list[j] {
            // 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
            for ; j >= 0 && deal < list[j]; j-- {
                list[j+1] = list[j] // 某数后移，给待排序留空位
            }
            list[j+1] = deal // 结束了，待排序的数插入空位
        }
    }
}

func main() {
    list := []int{5}
    InsertSort(list)
    fmt.Println(list)

    list1 := []int{5, 9}
    InsertSort(list1)
    fmt.Println(list1)

    list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    InsertSort(list2)
    fmt.Println(list2)
}