package main

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"encoding/base64"
	// 	"io/ioutil"
	"os"
)

// 定义接收数据的结构体
type QrCode struct {
    Url	string 
}


type msg struct {
	Name []byte
}


func ImagesToBase64() []byte {
	//读原图片
	ff, _ := os.Open("./policy_qrcode.png")
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	return []byte(sourcestring)
}


func main() {
    // 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	r.POST("/policy/url-qrcode", func(c *gin.Context) {
		// var form QrCode
		url := c.PostForm("Url")
		qrcode.WriteFile(url, qrcode.Medium, 256, "./policy_qrcode.png")

		data := msg{
			ImagesToBase64(),
		}
		c.JSON(http.StatusOK, data)
	})
	// 3.监听端口，默认在8080。Run("里面不指定端口号默认为8080")
	r.Run(":8005")

}
