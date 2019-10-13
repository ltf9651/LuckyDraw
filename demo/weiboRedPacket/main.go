// 微博红包

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"sync"
	"time"
)

//var packageList = make(map[uint32][]uint)
var packageList = new(sync.Map) // map加锁 抗并发 线程安全

type task struct {
	id       uint32
	callback chan uint
}

//packageList更新任务队列（代替锁）
//var chTasks = make(chan task)
// 单任务->多任务
const taskNum = 16

var chTaskList = make([]chan task, taskNum)

type lotteryController struct {
	Ctx iris.Context
}

//返回iris的Application
func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	for i := 0; i < taskNum; i++ {
		chTaskList[i] = make(chan task)
		go fetchPackageListMoney(chTaskList[i])
	}
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr("localhost:8080"))
}

// 返回所有红包地址
func (c *lotteryController) Get() map[uint32][2]int {
	result := make(map[uint32][2]int)
	/*for id, list := range packageList {
		var money int
		for _, v := range list {
			money += int(v)
		}
		result[id] = [2]int{len(list), money}
	}*/
	packageList.Range(func(key, value interface{}) bool {
		id := key.(uint32)
		list := value.([]uint)
		var money int
		for _, v := range list {
			money += int(v)
		}
		result[id] = [2]int{len(list), money}
		return true
	})
	return result
}

//发红包
func (c *lotteryController) GetSet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	money, errMoney := c.Ctx.URLParamFloat64("money")
	num, errNum := c.Ctx.URLParamInt("num")
	if errMoney != nil || errNum != nil || errUid != nil {
		return fmt.Sprintf("%d,%d,%d", errUid, errNum, errMoney)
	}
	moneyTotal := int(money * 100)
	leftNum := num
	leftMoney := moneyTotal

	//金额分配算法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rMax := 0.5 // 一次最多抢该红包总金额的50%
	list := make([]uint, num)
	for leftNum > 0 {
		if leftNum == 1 {
			//最后一个红包获得所有剩余金额
			list[num-1] = uint(leftMoney)
			break
		}
		if leftMoney == leftNum {
			// 剩下5分钱，5个红包，不进行分拆
			for i := num - leftNum; i < num; i++ {
				list[i] = 1
			}
			break
		}
		//随机
		rMoney := int(float64(leftMoney-leftNum) * rMax)
		m := r.Intn(rMoney)
		if m < 1 {
			m = 1
		}
		list[num-leftNum] = uint(m)
		leftMoney -= m
		leftNum--
	}

	//红包唯一id
	id := r.Uint32()
	//packageList[id] = list
	packageList.Store(id, list)
	//返回抢红包地址
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d", id, uid, num)
}

//抢红包
func (c *lotteryController) GetGet() string {
	id, errId := c.Ctx.URLParamFloat64("id")
	if errId != nil {
		return fmt.Sprintf("%d", errId)
	}

	//list, ok := packageList[uint32(id)]
	list1, ok := packageList.Load(uint32(id))
	list := list1.([]uint)
	if !ok || len(list) < 1 {
		// 红包为空或者抢失败
		return fmt.Sprintf("红包不存在")
	}

	//构造抢红包任务
	callback := make(chan uint)
	t := task{id: uint32(id), callback: callback}
	chTasks := chTaskList[int(id)%taskNum]
	chTasks <- t        //往Channel塞任务
	money := <-callback // 从Channel中接收数据

	// 分配随机数
	/*r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(list))
	money := list[i]
	//更新红包列表信息
	if len(list) > 1 {
		if i == len(list)-1 {
			//该红包是所有红包列表的最后一个红包
			//packageList[uint32(id)] = list[:i]
			packageList.Store(uint32(id), list[:i])
		} else if i == 0 {
			//packageList[uint32(id)] = list[1:]
			packageList.Store(uint32(id), list[1:])
		} else {
			//packageList[uint32(id)] = append(list[:i], list[i+1:]...)
			packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
		}
	} else {
		//该红包抢完，从红包列表中移除
		//delete(packageList, uint32(id))
		packageList.Delete(uint32(id))
	}
	*/
	return fmt.Sprintf("抢到：%d元\n", money)
}

// 消费抢红包任务
func fetchPackageListMoney(chTasks chan task) {
	for {
		t := <-chTasks
		id := t.id

		l, ok := packageList.Load(id)
		if ok && l != nil {
			list := l.([]uint)

			// 分配随机数
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			i := r.Intn(len(list))
			money := list[i]
			//更新红包列表信息
			if len(list) > 1 {
				if i == len(list)-1 {
					//该红包是所有红包列表的最后一个红包
					//packageList[uint32(id)] = list[:i]
					packageList.Store(uint32(id), list[:i])
				} else if i == 0 {
					//packageList[uint32(id)] = list[1:]
					packageList.Store(uint32(id), list[1:])
				} else {
					//packageList[uint32(id)] = append(list[:i], list[i+1:]...)
					packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
				}
			} else {
				//该红包抢完，从红包列表中移除
				//delete(packageList, uint32(id))
				packageList.Delete(uint32(id))
			}
			t.callback <- money // 回调，返回money
		} else {
			t.callback <- 0
		}
	}
}
