package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type AdvertRequest struct {
	Title   string `json:"title" binding:"required" msg:"Please input title" structs:"title"`
	Href    string `json:"href" binding:"required,url" msg:"Please input advertise href"` // binding中的url提供一种校验功能, 用于校验参数的数据内容是否是url
	Images  string `json:"images" binding:"required,url" msg:"Please input image href"`
	IsShown bool   `json:"is_shown" msg:"Please choose Shown Status" structs:"is_shown"` // bool值如果使用binding:"required", 传入的值必须为True, 否则被判为不满足条件
}

func main() {
	u1 := AdvertRequest{
		Title:   "xxxx",
		Href:    "xxxx",
		Images:  "xxxx",
		IsShown: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
