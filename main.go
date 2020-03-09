package main

import (
	"bloger/controller"
	"bloger/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:mysqlxjl12322@163.com@tcp(152.136.43.225:3306)/bloger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	//加载静态文件
	router.Static("/static/", "./static")
	//加载模板
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	_ = router.Run(":8001")
}