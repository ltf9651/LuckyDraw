package datasource

import (
	"LuckyDraw/project/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gpmgo/gopm/modules/log"
	"sync"
)

var masterInstance *xorm.Engine
var dbLock sync.Mutex

func InstanceDbMaster() *xorm.Engine {
	if masterInstance != nil {
		return masterInstance
	}

	dbLock.Lock() // 防止并发多个请求都去创建
	defer dbLock.Unlock()
	if masterInstance != nil {
		// unlock后再进行一次判断
		return masterInstance
	}

	return NewDbMaster()
}

func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@(tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database,
	)
	instace, err := xorm.NewEngine(conf.DriverName, sourcename)
	if err != nil {
		log.Fatal("error")
		return nil
	}
	instace.ShowSQL(true)
	masterInstance = instace
	return instace
}
