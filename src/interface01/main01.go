/*
接口的好处：(请注意，接口有且仅有方法定义，不允许有变量)
用一个结构体或其它对象来实现接口，我们到时传一个接口进一个方法，
接口会自动识别传入的对象是什么（前提是这个对象要实现这个接口）
接口会自动去找个对象找这个都西昂实现的接口里面的方法来执行
这，也就解释了为什么，接口的方法只有定义没有实现，就是未来动态化（多态和高内聚低耦合）
*/

package main

import (
	"fmt"
	"strconv"
)

// 声明一个接口
type Usb interface {
	Start(a int)
	Stop(b string)
}

type Phone struct {
}

type Camera struct {
}

func (p Phone) Start(a int) {
	o := strconv.Itoa(int(a))
	fmt.Println("手机" + o + "开始工作。。。")
}

func (p Phone) Stop(b string) {
	fmt.Println("手机" + b + "停止工作。。。")
}

//编写一个Working方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（必须实现接口内部所有方法）的对象就可以传入
func (p Camera) Working(usb Usb) {
	usb.Start(1)
	usb.Stop("lan")
}

func main() {
	//先创建结构体变量
	p := Phone{}
	c := Camera{}
	//关键点
	c.Working(p)

}
