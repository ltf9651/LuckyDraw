package controller

import (
	"LuckyDraw/project/common"
	"LuckyDraw/project/conf"
	"LuckyDraw/project/models"
	"LuckyDraw/project/services"
	"LuckyDraw/project/web/util"
)

func (c *IndexController) GetLucky() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["message"] = ""

	loginuser := comm.GetLoginUser(c.Ctx.Request())
	if loginuser == nil {
		rs["code"] = 101
		rs["message"] = "请先登录"
		return rs
	}

	// 用户分布式锁
	ok := util.LockLucky(loginuser.Uid)
	if ok {
		defer util.UnlockLucky(loginuser.Uid)
	} else {
		rs["code"] = 103
		rs["message"] = "操作频繁，稍后重试"
		return rs
	}

	userDayOK := c.CheckUserday(loginuser.Uid) // 验证今日次数
	if userDayOK {
		rs["code"] = 103
		rs["message"] = "今日次数已用完"
		return rs
	}

	ip := comm.ClientIP(c.Ctx.Request())
	ipDayNum := util.IncrIpLuckyNum(ip)
	if ipDayNum > conf.IpLimitMax {
		rs["code"] = 103
		rs["message"] = "同一IP今日次数已用完"
		return rs
	}

	// 开始抽奖
	prizeCode := comm.Random(10000)
	prizeGift := c.prize(prizeCode, false)
	if prizeGift == nil ||
		prizeGift.PrizeNum < 0 ||
		prizeGift.LeftNum <= 0 {
		rs["code"] = 205
		rs["message"] = "没中奖"
		return rs
	}

	rs["code"] = 200
	rs["msg"] = "恭喜中奖"
	rs["gift"] = prizeGift

	// 11 记录中奖记录
	result := models.LtResult{
		GiftId:     prizeGift.Id,
		GiftName:   prizeGift.Title,
		GiftType:   prizeGift.Gtype,
		Uid:        loginuser.Uid,
		Username:   loginuser.Username,
		PrizeCode:  prizeCode,
		GiftData:   prizeGift.Gdata,
		SysCreated: comm.NowUnix(),
		SysIp:      ip,
		SysStatus:  0,
	}
	_ := services.NewResultService().Create(&result)
	return rs
}
