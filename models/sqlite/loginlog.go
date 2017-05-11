package sqlite

import "time"

//登录日志数据操作
var LoginLogDb *LoginLog

func init() {
	LoginLogDb = NewLoginLog()
}

type LoginLog struct {
	Id        int    `json:"id" xorm:"not null unique pk autoincr INT(11)"`
	Username  string `json:"username" xorm:"not null default '' CHAR(30)"`
	LoginTime string `json:"login_time" xorm:"datetime"`
	LoginIp   string `json:"login_ip" xorm:"not null default '' char(15)"`
	Status    int    `json:"status" xorm:"not null default 0 index TINYINT(1)"`
	Info      string `json:"info" xorm:"not null default '' VARCHAR(66)"`
	Area      string `json:"area" xorm:"not null default '' VARCHAR(50)"`
	Country   string `json:"country" xorm:"not null default '' VARCHAR(50)"`
	Useragent string `josn:"useragent" xorm:"not null default 0 TEXT"`
}

func NewLoginLog() *LoginLog {
	return &LoginLog{}
}
func (l *LoginLog) RecordLogin(username string, loginip string, status int, info string, area string, country string, useragent string) error {
	l.LoginTime = time.Now().Format("2006-01-02 15:04:04")
	l.Username = username
	l.LoginIp = loginip
	l.Status = status
	l.Info = info
	l.Area = area
	l.Country = country
	l.Useragent = useragent
	_, err := x.InsertOne(l)
	return err
}
