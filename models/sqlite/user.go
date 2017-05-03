package sqlite

import (
	"hmcms/models"
	"time"
)

var (
	UserSql *User
)

func init() {
	UserSql = NewUser()
}

type User struct {
	Id            int       `json:"id" xorm:"not null unique pk autoincr INT(11)"`
	Username      string    `json:"username" xorm:"not null unique default '' index CHAR(30)"`
	Password      string    `json:"password" xorm:"not null default '' CHAR(32)"`
	Nickname      string    `json:"nickname" xorm:"not null default '' VARCHAR(50)"`
	Email         string    `json:"email" xorm:"not null default '' VARCHAR(80)"`
	LastLoginTime time.Time `json:"last_login_time" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateTime    time.Time `json:"create_time" xorm:"TIMESTAMP"`
	LastLoginIp   string    `json:"last_login_ip" xorm:"not null default '' CHAR(15)"`
	CreateIp      string    `json:"create_ip" xorm:"not null default '' CHAR(15)"`
	Remake        string    `json:"remake" xorm:"not null default '' VARCHAR(255)"`
	Status        int       `json:"status" xorm:"not null default 0 index TINYINT(1)"`
}

func NewUser() *User {
	return &User{}
}
func (u *User) LoginUser(username string, password string) (bool, error) {
	password = models.Sha1([]byte(models.Md5([]byte(password))))
	has, err := x.Where("username = ? AND password = ?", username, password).Get(u)
	return has, err
}
