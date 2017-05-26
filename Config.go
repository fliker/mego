package config

type Config struct {
	Index           string
	Controller      string
	Action          string
	SessionType     int      //0,不使用Session 1,文件,2内存
	MemFreeInterval int      //单位，秒
	SessionTimeOut  int      //Session超时时间，单位分钟
	SessionLocation string   //Session文件存放位置，为1时才有效
	StaticFile      []string //静态文件
	StaticDir       []string //静态目录
	Location        string
}

func Default() *Config {
	c := &Config{}
	c.Index = "home/login.api"
	c.Controller = "index"
	c.Action = "do"
	c.SessionType = 0
	c.MemFreeInterval = 300
	c.SessionTimeOut = 35
	c.SessionLocation = "data/session"
	c.StaticDir = []string{"public/"}
	//c.StaticFile = []string{}
	c.Location = "root:root@tcp(127.0.0.1:3306)/caiji?charset=utf8"
	return c
}
