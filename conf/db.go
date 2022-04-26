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

// 系统中所有mysql主库 root:root@tcp(192.168.56.10:3306)/lottery?charset=utf-8
var DbMasterList = []DbConfig{
	{
		Host:      "192.168.56.10",
		Port:      3306,
		User:      "root",
		Pwd:       "root",
		Database:  "lottery",
		IsRunning: true,
	},
}

var DbMaster DbConfig = DbMasterList[0]
