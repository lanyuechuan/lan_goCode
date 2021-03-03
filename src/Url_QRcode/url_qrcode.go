package main

import (
	"net/http"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"unsafe"
	"os"
	"bytes"
	"io"
	"io/ioutil"
)

// 定义接收数据的结构体
type QrCode struct {
    Url	string `form:"url" json:"url" uri:"url" xml:"url" binding:"required"`
}


func ImagesToBase64(insurance string) string {
	//读原图片
	fileName := ""

	if insurance == "PICC" {
		fileName = "./picc_policy_qrcode.png"
	}else {
		fileName = "./gpic_policy_qrcode.png"
	}

	ff, _ := os.Open(fileName)
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	baseImage := []byte(sourcestring)
	// *(*string)(unsafe.Pointer(&baseImage)) --->切片转字符串
	return "data:image/png;base64," + *(*string)(unsafe.Pointer(&baseImage))

}


func getImg(url string) (n int64, err error) {
	var name string = "gpic_policy_qrcode.png"
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return

}

func main() {
    // 1.创建路由
	r := gin.Default()

	//该接口把人保的支付链接转换为二维码图片，然后返回base64（源文件是支付链接）
	r.POST("/picc/policy/url-qrcode", func(c *gin.Context) {
		var form QrCode
		url := c.PostForm("url")
		if err := c.Bind(&form); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error(), "status":"400"})
			return
		}
		qrcode.WriteFile(url, qrcode.Medium, 256, "./picc_policy_qrcode.png")

		url_code := ImagesToBase64("PICC")
		c.JSON(http.StatusOK, gin.H{"url_code":url_code, "status":"200"})

		
	})


	// 该接口是把国寿的支付二维码图片下载下来，然后转换为base64返回（源文件是支付二维码图片链接）
	r.POST("/gpic/policy/url-qrcode", func(c *gin.Context) {
		var form QrCode
		url := c.PostForm("url")
		if err := c.Bind(&form); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error(), "status":"400"})
			return
		}
		getImg(url)

		url_code := ImagesToBase64("GPIC")
		c.JSON(http.StatusOK, gin.H{"url_code":url_code, "status":"200"})
	})

	// 3.监听端口，默认在8080。Run("里面不指定端口号默认为8080")
	r.Run(":8003")

}
