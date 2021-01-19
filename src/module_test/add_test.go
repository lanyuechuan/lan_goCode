/*
注意：测试用例文件名必须用_test.go结尾，测试函数必须用Test+函数名组成，并且函数名必须大写(大写表示可以外部调用)
形参必须是*testing.T，一个测试用例文件中可以有多个测试函数
测试运行命令，go test -v ,-v表示运行下正确或错误都会输出日志，没加就只有错误时才有日志
t.Fatalf输出错误信息并退出程序
t.Logf输出指定日志
PASS表示测试用例成功，FAIL表示测试用例失败
测试单个文件，一定要带上被测试的源文件
测试单个方法则带上测试函数名
*/
package main


import (
	_ "fmt"
	"testing"
)

func TestAddNum(t *testing.T) {
	//调用AddNum
	res := AddNum(10)
	if res != 55{
		// fmt.Println("执行错误")
		t.Fatalf("执行错误")
	}
	//如果正确，输出日志
	t.Logf("执行正确")
}