package lib


import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)


var Gorm *gorm.DB
func initDB() {
	Gorm=gormDB()
}

func gormDB() *gorm.DB {
	dsn := "user:password@tcp(192.168.x.x:3306)/cmdb?charset=utf8mb4&parseTime=True&loc=Local"

	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	mysqlDB,err:= db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetMaxIdleConns(5)
	return db
}
