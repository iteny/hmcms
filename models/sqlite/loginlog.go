package sqlite

import "time"

//登录日志数据操作
var LoginLogDb *LoginLog

func init() {
	LoginLogDb = NewLoginLog()
}

type LoginLog struct {
	Id        int       `json:"id" xorm:"not null unique pk autoincr INT(11)"`
	Username  string    `json:"username" xorm:"not null default '' CHAR(30)"`
	LoginTime time.Time `json:"login_time" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	LoginIp   string    `json:"login_ip" xorm:"not null default '' char(15)"`
	Status    int       `json:"status" xorm:"not null default 0 index TINYINT(1)"`
	Info      string    `json:"info" xorm:"not null default '' VARCHAR(66)"`
	Area      string    `json:"area" xorm:"not null default '' VARCHAR(50)"`
	Country   string    `json:"country" xorm:"not null default '' VARCHAR(50)"`
	Useragent string    `josn:"useragent" xorm:"not null default 0 TEXT"`
}

func NewLoginLog() *LoginLog {
	return &LoginLog{}
}
