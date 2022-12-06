package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
@author:君
@date:2022-12-05
@note:
*/
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main.recover: ", r)
		}
	}()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if r.Run() != nil {
		panic("gin启动失败")
	} // 监听并在 0.0.0.0:8080 上启动服务
}
