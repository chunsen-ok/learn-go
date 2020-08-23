package main

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "play-gin/docs"
	"play-gin/db"
)

// @title 我的Swagger API
// @version 0.0.1
// @license.name LGPL 3.0

func main() {
	d := db.Setup()

	g := gin.Default()
	g.GET("/home",func(c *gin.Context){

		id, err := strconv.Atoi(c.Query("id"))
		if err == nil {
			c.JSON(200,gin.H{
				"message": db.Query(d,int64(id)),
			})
		} else {
			fmt.Println(err)
		}
	})
	g.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.Run("127.0.0.1:9001")
}

// Test 简单打印信息
// @summary 简要叙述这个API的功能
// @description 这个API是用来获取用户信息的，假如可以用的话:)
// @tags 基本API
// @accept plain
// @produce json
// @param id query integer true "用户ID"
// @success 200 body gin.H "用户详情"
// @router /home [GET]
func Test() {
	fmt.Println("This is test")
}
