package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float32  `json:"price"`
	Actors []string `json:"actors"`
}


func main() {
	movie := Movie{"喜剧之王", 2000, 16.8, []string{"lan","yue"}}
	//编码过程，结构体转json
	jsonByte, err := json.Marshal(movie)
	jsonStr := string(jsonByte)
	if err != nil{
		fmt.Println("失败")
		return
	}
		
	fmt.Printf(jsonStr) //是Printf

	//解码过程，json转结构体
	my_movie := Movie{} // 这个my_movie相当于是一个解析模板，按这个结构体模板来解析
	err = json.Unmarshal([]byte(jsonStr),&my_movie)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(my_movie)
}