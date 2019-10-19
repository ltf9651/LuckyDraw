package models

type LuckyRecord struct {
	Id         int    `xorm:"not null pk INT(11)"`
	GiftId     int    `xorm:"INT(11)"`
	GiftName   string `xorm:"VARCHAR(255)"`
	GiftType   int    `xorm:"INT(11)"`
	Uid        int    `xorm:"INT(11)"`
	Username   string `xorm:"VARCHAR(255)"`
	PrizeCode  int    `xorm:"INT(11)"`
	GiftData   string `xorm:"VARCHAR(255)"`
	SysCreated int    `xorm:"INT(11)"`
	SysIp      string `xorm:"VARCHAR(255)"`
	SysStatus  int    `xorm:"SMALLINT(6)"`
}
