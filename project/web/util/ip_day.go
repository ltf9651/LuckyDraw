package util

import (
	comm "LuckyDraw/project/common"
	"LuckyDraw/project/datasource"
	"fmt"
)

func IncrIpLuckyNum(strIp string) int64 {
	ip := comm.Ip4toInt(strIp)
	i := ip % 2
	key := fmt.Sprintf("day_ips_%d", i)
	cacheObj := datasource.InstanceCache()
	rs, _ := cacheObj.Do("HINCRBY", key, ip, 1)
	return rs.(int64)
}
