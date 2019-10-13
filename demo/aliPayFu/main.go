// 支付宝集福

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type gift struct {
	id      int
	name    string
	rate    int //中奖概率
	rateMin int
	rateMax int
}

const rateMax = 20

type lotterController struct {
	Ctx iris.Context
}

var logger *log.Logger

func initLog() {
	f, _ := os.Create("/log.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newGift() *[5]gift {
	giftList := new([5]gift)
	g1 := gift{
		id:      1,
		name:    "福卡1",
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[0] = g1
	g2 := gift{
		id:      2,
		name:    "福卡2",
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[1] = g2
	g3 := gift{
		id:      3,
		name:    "福卡3",
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[2] = g3
	g4 := gift{
		id:      4,
		name:    "福卡4",
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[3] = g4
	g5 := gift{
		id:      5,
		name:    "福卡5",
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[4] = g5
	return giftList
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotterController{})
	initLog()
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr("localhost:8080"))
}

func setGiftRate(rate string) *[5]gift {
	// giftList每次请求的时候重新生成，不存在线程安全问题
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rates)

	rateStart := 0
	for i := range giftList {
		grate := 0
		if i < ratesLen {
			grate, _ = strconv.Atoi(rates[i])
		}
		giftList[i].rate = grate
		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + grate
		if giftList[i].rateMax >= rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}
	fmt.Printf("giftList = %v\n", giftList)

	return giftList
}

func (c *lotterController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "5,4,3,2,1")
	giftList := setGiftRate(rate)
	return fmt.Sprintf("%v\n", giftList)
}

// 抽奖
func (c *lotterController) GetLucky() map[string]interface{} {
	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	sendData := ""
	rate := c.Ctx.URLParamDefault("rate", "5,4,3,2,1")
	giftList := setGiftRate(rate)
	for _, data := range giftList {
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			//中奖
			sendData = data.name
			ok = true
		}
		if ok {
			saveLuckyData(code, data.id, data.name, sendData)
			result["success"] = ok
			result["id"] = data.id
			result["name"] = data.name
			break
		}
	}
	return result
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return code
}

func saveLuckyData(code int32, id int, name string, sendData string) {
	logger.Printf("luckyInfo : %d %d %s %s %s %d", code, id, name, sendData)
}
