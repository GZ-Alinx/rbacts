package lib

import (
	"github.com/gin-gonic/gin"
)



// 检查是否登录
func CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Header.Get("token")==""{
			context.AbortWithStatusJSON(400, gin.H{"message":"token required"})
		}else {
			context.Set("user_name", context.Request.Header.Get("token"))
			context.Next()
		}
	}
}


func RBAC() gin.HandlerFunc {
	// 模型和策略进行持久化方式
	//e := casbin.NewEnforcer("resources/model.conf", "resources/p.csv")
	//adapter,err := gormadapter.NewAdapterByDB(Gorm)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//e,err := casbin.NewEnforcer("resources/model.conf", adapter)
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//err = e.LoadPolicy()
	//if err!=nil {
	//	log.Fatal(err)
	//}
	return func(context *gin.Context) {
		user,_:= context.Get("user_name")
		access,err := E.Enforce(user,context.Request.RequestURI, context.Request.Method)
		if err != nil || !access{
			context.AbortWithStatusJSON(403, gin.H{"message":"forbidden"})
		}else{
			context.Next()
		}
	}
}


func Middlewares() (fs []gin.HandlerFunc) {
	fs = append(fs, CheckLogin(),RBAC())
	return
}