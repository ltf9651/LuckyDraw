package models

type UserDailyRecord struct {
	Id         int `xorm:"not null pk INT(11)"`
	Uid        int `xorm:"INT(11)"`
	Day        int `xorm:"INT(11)"`
	Num        int `xorm:"INT(11)"`
	SysCreated int `xorm:"INT(11)"`
	SysUpdated int `xorm:"INT(11)"`
}
