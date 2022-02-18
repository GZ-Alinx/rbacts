package lib

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var E *casbin.Enforcer

func init() {
	initDB()
	adapter,err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatal(err)
	}
	e, err := casbin.NewEnforcer("resources/model.conf", adapter)
	if err != nil {
		log.Fatal(err)
	}
	err =  e.LoadPolicy()
	if err!=nil {
		log.Fatal(err)
	}
	E=e
	initPolicy()
}


// 从我们的库里初始化策略数据
func initPolicy() {
	E.AddPolicy("member", "/depts", "GET")
	E.AddPolicy("administrator", "/depts", "POST")
	E.AddRoleForUser("lixiong","member")
}
