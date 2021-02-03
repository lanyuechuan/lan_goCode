package main

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
  )

type Gift struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Statement string `json:"statement" form:"statement"`
	Status string `json:"status" form:"status"`
	User_id string `json:"user_id" form:"user_id"`
	Organization string `json:"organization" form:"organization"`
}


func main() {
	db, err := gorm.Open("mysql", "devel_test:develtest20200402@tcp(10.252.252.18:3306)/qdtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println(err)
		}
	defer db.Close()
	router := gin.Default()
	router.GET("/order-gift", func(c *gin.Context) {
		rs, _ := getRows()
		c.JSON(http.StatusOK, gin.H{
			"list": rs,
		})
	})
	router.Run(":8801")
}

//查询所有记录
func getRows() (gifts []Gift, err error) {
	db, err := gorm.Open("mysql", "devel_test:develtest20200402@tcp(10.252.252.18:3306)/qdtest?charset=utf8&parseTime=True&loc=Local")
	rows, err := db.Query("select id,name,statement,status,user_id,organization from order_gift")
	for rows.Next(){
		gift := Gift{}
		err := rows.Scan(&gift.Id, &gift.Name, &gift.Statement, &gift.Status, &gift.User_id, &gift.Organization)
		if err != nil {
			fmt.Println(err)
		}
		gifts = append(gifts, gift)
	}
	rows.Close()
	return
}