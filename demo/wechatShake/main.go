package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var mu sync.Mutex

//奖品类型： 枚举 iota 从 0 开始
const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 不同券
	giftTypeCouponFix        // 相同券
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

type gift struct {
	id       int
	name     string
	pic      string
	link     string
	giftType int
	data     string
	dataList []string // 奖品数据集合
	total    int
	left     int  // 库存
	inUse    bool // 是否使用中
	rate     int  //中奖概率
	rateMin  int
	rateMax  int
}

// 最大中奖号码
const rateMax = 10000

var logger *log.Logger

var giftList []*gift

type lotterController struct {
	Ctx iris.Context
}

func initLog() {
	f, _ := os.Create("/log.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func initGift() {
	giftList = make([]*gift, 5)
	g1 := gift{
		id:       1,
		name:     "iphone 11",
		pic:      "",
		link:     "",
		giftType: giftTypeRealLarge,
		data:     "",
		dataList: nil,
		total:    1000,
		left:     1000,
		inUse:    true,
		rate:     100,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[0] = &g1
	g2 := gift{
		id:       2,
		name:     "手机壳",
		pic:      "",
		link:     "",
		giftType: giftTypeRealSmall,
		data:     "",
		dataList: nil,
		total:    20,
		left:     20,
		inUse:    true,
		rate:     500,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[1] = &g2
	g3 := gift{
		id:       3,
		name:     "优惠券",
		pic:      "",
		link:     "",
		giftType: giftTypeCouponFix,
		data:     "coupon-2019",
		dataList: nil,
		total:    20,
		left:     20,
		inUse:    true,
		rate:     5000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[2] = &g3
	g4 := gift{
		id:       4,
		name:     "直降优惠券",
		pic:      "",
		link:     "",
		giftType: giftTypeCoupon,
		data:     "",
		dataList: []string{"c_1", "c_2", "c_3"},
		total:    20,
		left:     20,
		inUse:    true,
		rate:     2000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[3] = &g4
	g5 := gift{
		id:       5,
		name:     "Q币",
		pic:      "",
		link:     "",
		giftType: giftTypeCoupon,
		data:     "10金币",
		dataList: nil,
		total:    20,
		left:     20,
		inUse:    true,
		rate:     5000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[4] = &g5

	// 中奖区间数据
	rateStart := 0
	for _, data := range giftList {
		if !data.inUse {
			continue
		}
		data.rateMin = rateStart
		data.rateMax = rateStart + data.rate
		if data.rateMax >= rateMax {
			data.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotterController{})
	initLog()
	initGift()
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr("localhost:8080"))
}

// 奖品数量信息
func (c *lotterController) Get() string {
	count := 0
	total := 0
	for _, data := range giftList {
		if data.inUse && (data.total == 0 || (data.total > 0 && data.left > 0)) {
			count++
			total += data.left
		}
	}
	return fmt.Sprintf("当前有效奖品种类数量：%d，总量：%d\n", count, total)
}

// 抽奖
func (c *lotterController) GetLucky() map[string]interface{} {
	mu.Lock() // 防止超卖
	defer mu.Unlock()

	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	sendData := ""
	for _, data := range giftList {
		if !data.inUse || (data.total > 0 && data.left <= 0) {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			//中奖
			switch data.giftType {
			case giftTypeCoin:
				ok, sendData = sendCoin(data)
			case giftTypeCoupon:
			case giftTypeCouponFix:
			case giftTypeRealLarge:
			case giftTypeRealSmall:
			}
		}
		if ok {
			saveLuckyData(code, data.id, data.name, data.link, sendData, data.left)
			result["success"] = ok
			result["id"] = data.id

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

func sendCoin(data *gift) (bool, string) {
	if data.total == 0 {
		// 无限量
		return true, data.data
	} else if data.left > 0 {
		data.left -= 1
		return true, data.data
	}
	return false, "奖品已发完"
}

func saveLuckyData(code int32, id int, name string, link string, sendData string, left int) {
	logger.Printf("luckyInfo : %d %d %s %s %s %d", code, id, name, link, sendData, left)
}
