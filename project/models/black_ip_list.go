package models

type BlackIpList struct {
	Id         int    `xorm:"not null pk INT(11)"`
	Ip         string `xorm:"VARCHAR(255)"`
	Blacktime  int    `xorm:"INT(11)"`
	SysCreated int    `xorm:"INT(11)"`
	SysUpdated int    `xorm:"INT(11)"`
}
