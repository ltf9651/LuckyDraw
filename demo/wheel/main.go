//大转盘

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
)

type lotterController struct {
	Ctx iris.Context
}

type Prate struct {
	Rate  int
	Total int
	CodeA int
	CodeB int
	Left  *int32
}

var prizeList = []string{
	"一等奖",
	"二等奖",
	"三等奖",
	"",
}

var left = int32(1000)

var rateList = []Prate{
	{1, 1, 0, 0, &left},
	{14, 1, 0, 0, &left},
	{1454, 1, 0, 0, &left},
	{100, 1, 999, 99999, &left},
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

func (c *lotterController) Get() string {
	//c.Ctx.Header("Content-type", "text/html")
	return fmt.Sprintf("奖品列表：\n %s", strings.Join(prizeList, "\n"))
}

func (c *lotterController) GetDenug() string {
	return fmt.Sprintf("抽奖概率 : %\n", rateList)
}

func (c *lotterController) GetPrize() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	code := r.Intn(1000)

	var myPrize string
	var prizeRate *Prate

	for i, prize := range prizeList {
		rate := &rateList[i]
		if code >= rate.CodeA && code <= rate.CodeB {
			myPrize = prize
			prizeRate = rate
			break //中奖
		}
	}

	if myPrize == "" {
		return "没中"
	}
	//发奖
	if prizeRate.Total == 0 {
		//无限制奖品数量
		return myPrize
	} else if *prizeRate.Left > 0 {
		//prizeRate.Left--

		//原子操作，性能高于互斥锁
		left := atomic.AddInt32(prizeRate.Left, -1)
		if left >= 0 {
			return myPrize
		}
	}
	return "没中"
}
