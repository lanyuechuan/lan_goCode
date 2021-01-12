/*
主要内容:
map是无序的，那怎么实现map的有序输出(有序肯定是针对键的大小升序或降序)
即新建一个切片把键装入切片中再排序

golang给了一套专门的切片排序，要先引入sort包，默认升序，有好几种，给字符串，浮点型，整型分别是Strings，Float64s,Ints
案例：
s := []string{"z","f","c","d","e","y","g"}
sort.Strings(s) // 输出[d e f g g y z]   
sort.Sort(sort.Reverse(sort.StringSlice(s))) // 降序比较麻烦，同理IntSlice和Float64Slice
*/

package main

import (
    "fmt"
    "sort"
)

func main() {
    map1 := make(map[int]string, 5)
    map1[1] = "www.topgoer.com"
    map1[2] = "rpc.topgoer.com"
    map1[5] = "ceshi"
    map1[3] = "xiaohong"
    map1[4] = "xiaohuang"
    sli := []int{}
    for k, _ := range map1 {
        sli = append(sli, k)
    }
    sort.Ints(sli)
    for i := 0; i < len(map1); i++ {
        fmt.Println(map1[sli[i]])
    }
}