package main

import (
	"github.com/gin-gonic/gin"
  )


func sayHello(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello", sayHello)

	// 启动服务
	router.Run()
}