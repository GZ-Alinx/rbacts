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
	// 模型和策略进行持久化方式
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
	//E.AddPolicy("member", "/depts", "GET")
	//E.AddPolicy("administrator", "/depts", "POST")
	//E.AddRoleForUser("lixiong","member")

	/////// 角色初始化
	m := make([]*RoleRel,0)
	GetRoles(0,&m,"")// 获取角色对应
	for _,r := range m {
		_,err := E.AddRoleForUser(r.PRole,r.Role)
		if err!=nil {
			log.Fatal(err)
		}
	}

	////// 初始化 用户角色
	userRoles := GetUserRoles()
	for _,ur := range userRoles{
		_,err := E.AddRoleForUser(ur.RoleName,ur.RoleName)
		if err != nil {
			log.Fatal(err)
		}
	}

	////// 初始化 路由角色
	routerRoles := GetRouterRoles()
	for _,rr := range routerRoles {
		_,err := E.AddPolicy(rr.RoleName, rr.RouterUri, rr.RouterMethod)
		if err != nil {
			log.Fatal(err)
		}
	}
}
