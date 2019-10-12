package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync"
	"time"
)

//用户列表
var userList []string
var mu sync.Mutex // 加锁 保证线程安全

type lotteryController struct {
	Ctx iris.Context
}

//返回iris的Application
func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	userList = []string{}
	mu = sync.Mutex{}

	app.Run(iris.Addr("localhost:8080"))
}

func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前参与抽奖人数：%d\n", count)
}

// POST localhost:8080/import
// params: users
// 导入用户
func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")

	mu.Lock()
	defer mu.Unlock()

	count1 := len(userList)
	for _, u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}

	count2 := len(userList)
	return fmt.Sprintf("当前参与人数：%d，成功导入人数：%d\n", count2, count2-count1)
}

// POST localhost:8080/lucky
// 抽奖
func (c *lotteryController) GetLucky() string {
	mu.Lock()
	defer mu.Unlock()

	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else if count == 1 {
		user := userList[0]
		return fmt.Sprintf("中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else {
		return fmt.Sprintf("用户列表为空")
	}
}
