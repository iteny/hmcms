package sqlite

//登录日志数据操作
var AuthRuleDb *AuthRule

func init() {
	AuthRuleDb = NewAuthRule()
}

type AuthRule struct {
	Id       int         `json:"id" xorm:"integer(11) notnull unique pk autoincr"`
	Url      string      `json:"url" xorm:"text(80) notnull default ''"`
	Name     string      `json:"name" xorm:"text(20) notnull default ''"`
	Pid      int         `json:"pid" xorm:"integer(6) notnull default 0 index"`
	Isshow   int         `json:"isshow" xorm:"integer(1) notnull default 0 index"`
	Sort     int         `json:"sort" xorm:"integer(6) notnull default 0 index"`
	Icon     string      `json:"icon" xorm:"text(50) notnull default ''"`
	Children []*AuthRule `json:"children"`
}

func NewAuthRule() *AuthRule {
	return &AuthRule{}
}

//获取一级菜单
func (a *AuthRule) GetOneMenu() (ar []AuthRule, err error) {
	err = x.Where("pid = ? AND isshow = ?", 0, 1).Find(&ar)
	return ar, err
}

//获取二级菜单
func (a *AuthRule) GetTwoMenu(pid int) (ar []*AuthRule, err error) {
	err = x.Asc("sort").Cols("id,url,name,icon").Where("pid= ? AND isshow= ?", pid, 1).Find(&ar)
	return ar, err
}
