package main

import (
	"github.com/gin-gonic/gin"
	"rbacts/lib"
)



// Casbin 管理api的使用  https://casbin.org/docs/zh-CN/management-api
// 			RBAC管理API  https://casbin.org/docs/zh-CN/rbac-api
func main() {
	r := gin.New()
	r.Use(lib.Middlewares()...)
	r.GET("/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{"result":"部门列表"})
	})
	r.POST("/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{"result":"批量修改部门列表"})
	})

	r.Run(":8081")
}


