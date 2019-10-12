/*
1. 刮刮乐
2. 双色球
*/
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"time"
)

type lotterController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotterController{})
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr("localhost:8080"))
}

// 刮刮乐
func (c *lotterController) Get() string {
	var code int
	var prize string

	seed := time.Now().UnixNano()
	code = rand.New(rand.NewSource(seed)).Intn(10)
	switch {
	case code == 1:
		prize = "一等奖"
	case code == 2:
		prize = "二等奖"
	case code == 3:
		prize = "三等奖"
	default:
		prize = "没中奖"
	}
	return fmt.Sprintf(prize)
}

// 双色球
func (c *lotterController) GetPrize() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	var prize [7]int
	//随机生成七位数
	for i := 0; i < 6; i++ {
		prize[i] = r.Intn(33) + 1 // [1,33)
	}
	prize[6] = r.Intn(16) + 1

	return fmt.Sprintf("今日开奖号码： %v\n", prize)
}
