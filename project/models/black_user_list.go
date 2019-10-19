package models

type BlackUserList struct {
	Id         int    `xorm:"not null pk INT(11)"`
	Username   string `xorm:"VARCHAR(255)"`
	Blacktime  int    `xorm:"INT(11)"`
	Realname   string `xorm:"VARCHAR(255)"`
	Mobile     string `xorm:"VARCHAR(255)"`
	Address    string `xorm:"VARCHAR(255)"`
	SysCreated int    `xorm:"INT(11)"`
	SysUpdated int    `xorm:"INT(11)"`
	SysIp      string `xorm:"VARCHAR(255)"`
}
