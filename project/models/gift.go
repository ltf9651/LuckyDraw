package models

type Gift struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	GiftId     int    `xorm:"INT(11)"`
	Code       string `xorm:"VARCHAR(255)"`
	SysCreated int    `xorm:"INT(11)"`
	SysUpdated int    `xorm:"INT(11)"`
	SysStatus  int    `xorm:"SMALLINT(6)"`
}
