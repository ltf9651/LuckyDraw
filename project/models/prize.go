package models

type Prize struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	Title        string `xorm:"VARCHAR(255)"`
	PrizeNum     int    `xorm:"INT(11)"`
	LeftNum      int    `xorm:"INT(11)"`
	PrizeCode    string `xorm:"comment('概率') VARCHAR(50)"`
	Img          string `xorm:"VARCHAR(255)"`
	Displayorder int    `xorm:"comment('位置序号') INT(11)"`
	Gtype        int    `xorm:"comment('奖品类型') INT(11)"`
	Gdata        string `xorm:"comment('拓展数据') VARCHAR(255)"`
	TimeBegin    int    `xorm:"INT(11)"`
	TimeEnd      int    `xorm:"INT(11)"`
	PrizeData    string `xorm:"comment('发奖计划') MEDIUMTEXT"`
	PrizeBegin   int    `xorm:"INT(11)"`
	PrizeEnd     int    `xorm:"INT(11)"`
	SysStatus    int    `xorm:"SMALLINT(6)"`
	SysCreated   int    `xorm:"INT(11)"`
	SysUpdated   int    `xorm:"INT(11)"`
	SysIp        string `xorm:"VARCHAR(50)"`
}
