package entities

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	orm, err1 := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err1)
	err1 = orm.Sync2(new(UserInfo))
	checkErr(err1)
	engine = orm
	
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
