package conf

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool
}

var DbMasterList = []DbConfig{
	{
		Host:      "localhost",
		Port:      3306,
		User:      "root",
		Pwd:       "111",
		Database:  "prize",
		IsRunning: true,
	},
}

var DbMaster = DbMasterList[0]
