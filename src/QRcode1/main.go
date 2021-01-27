package main

import (
	"net/http"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"unsafe"
	"os"
)

// 定义接收数据的结构体
type QrCode struct {
    Url	string `form:"url" json:"url" uri:"url" xml:"url" binding:"required"`
}


func ImagesToBase64() string {
	//读原图片
	ff, _ := os.Open("./policy_qrcode.png")
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	baseImage := []byte(sourcestring)
	// *(*string)(unsafe.Pointer(&baseImage)) --->切片转字符串
	return "data:image/png;base64," + *(*string)(unsafe.Pointer(&baseImage))

}


func main() {
    // 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	r.POST("/policy/url-qrcode", func(c *gin.Context) {
		var form QrCode
		url := c.PostForm("url")
		if err := c.Bind(&form); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error(), "status":"400"})
			return
		}
		qrcode.WriteFile(url, qrcode.Medium, 256, "./policy_qrcode.png")

		url_code := ImagesToBase64()
		c.JSON(http.StatusOK, gin.H{"url_code":url_code, "status":"200"})
	})
	// 3.监听端口，默认在8080。Run("里面不指定端口号默认为8080")
	r.Run(":8005")

}
