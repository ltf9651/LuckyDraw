package conf

type RdsConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	IsRunning bool
}

var RdsCacheList = []RdsConfig{
	{
		Host:      "localhost",
		Port:      6379,
		User:      "root",
		Pwd:       "111",
		IsRunning: true,
	},
}

var RdsCache = RdsCacheList[0]
