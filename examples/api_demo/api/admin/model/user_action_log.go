package model

type UserActionLog struct {
	Id            uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	UserId        uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:'0'" json:"user_id"`              // 用户id
	Count         uint   `gorm:"column:count;type:int(10) unsigned;default:'0'" json:"count"`                     // 访问次数
	LastVisitTime uint   `gorm:"column:last_visit_time;type:int(10) unsigned;default:'0'" json:"last_visit_time"` // 最后访问时间
	Object        string `gorm:"column:object;type:varchar(100);default:''" json:"object"`                        // 访问对象的id,格式:不带前缀的表名+id;如posts1表示xx_posts表里id为1的记录
	Action        string `gorm:"column:action;type:varchar(50);default:''" json:"action"`                         // 操作名称;格式:应用名+控制器+操作名,也可自己定义格式只要不发生冲突且惟一;
	Ip            string `gorm:"column:ip;type:varchar(15);default:''" json:"ip"`                                 // 用户ip
}

//get real primary key name
func (userActionLog *UserActionLog) GetKey() string {
	return "id"
}

//get primary key in model
func (userActionLog *UserActionLog) GetKeyProperty() uint64 {
	return userActionLog.Id
}

//set primary key
func (userActionLog *UserActionLog) SetKeyProperty(id uint64) {
	userActionLog.Id = id
}

//get real table name
func (userActionLog *UserActionLog) TableName() string {
	return "cmf_user_action_log"
}
