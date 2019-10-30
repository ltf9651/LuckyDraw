package controller

import (
	"LuckyDraw/project/conf"
	"LuckyDraw/project/models"
)

func (c *IndexController) CheckUserday(uid int) bool {
	return true
}

func (c *IndexController) prize(prizeCode int, limitBlock bool) *models.ObjGiftPrize {
	var prizeGift *models.ObjGiftPrize
	giftList := c.ServiceGift.GetAllUse(true)
	for _, gift := range giftList {
		if gift.PrizeCodeA <= prizeCode &&
			gift.PrizeCodeB >= prizeCode {
			//中奖
			if !limitBlock || gift.Gtype < conf.GtypeGiftSmall {
				prizeGift = &gift
				break
			}
		}
	}
	return prizeGift
}
