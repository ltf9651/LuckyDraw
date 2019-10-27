package controller

import (
	comm "LuckyDraw/project/common"
	"LuckyDraw/project/conf"
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

	return rs
}
