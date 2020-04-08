package conf

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var DbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   8306,
	User:   "root",
	Pwd:    "123456",
	DbName: "test",
}
