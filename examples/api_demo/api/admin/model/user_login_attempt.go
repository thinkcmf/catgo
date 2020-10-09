package model

type UserLoginAttempt struct {
	Id            uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	LoginAttempts uint   `gorm:"column:login_attempts;type:int(10) unsigned;default:'0'" json:"login_attempts"` // 尝试次数
	AttemptTime   uint   `gorm:"column:attempt_time;type:int(10) unsigned;default:'0'" json:"attempt_time"`     // 尝试登录时间
	LockedTime    uint   `gorm:"column:locked_time;type:int(10) unsigned;default:'0'" json:"locked_time"`       // 锁定时间
	Ip            string `gorm:"column:ip;type:varchar(15);default:''" json:"ip"`                               // 用户 ip
	Account       string `gorm:"column:account;type:varchar(100);default:''" json:"account"`                    // 用户账号,手机号,邮箱或用户名
}

//get real primary key name
func (userLoginAttempt *UserLoginAttempt) GetKey() string {
	return "id"
}

//get primary key in model
func (userLoginAttempt *UserLoginAttempt) GetKeyProperty() uint64 {
	return userLoginAttempt.Id
}

//set primary key
func (userLoginAttempt *UserLoginAttempt) SetKeyProperty(id uint64) {
	userLoginAttempt.Id = id
}

//get real table name
func (userLoginAttempt *UserLoginAttempt) TableName() string {
	return "cmf_user_login_attempt"
}
