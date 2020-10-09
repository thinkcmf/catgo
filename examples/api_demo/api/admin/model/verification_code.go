package model

type VerificationCode struct {
	Id         uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`                            // 表id
	Count      uint   `gorm:"column:count;type:int(10) unsigned;default:'0'" json:"count"`             // 当天已经发送成功的次数
	SendTime   uint   `gorm:"column:send_time;type:int(10) unsigned;default:'0'" json:"send_time"`     // 最后发送成功时间
	ExpireTime uint   `gorm:"column:expire_time;type:int(10) unsigned;default:'0'" json:"expire_time"` // 验证码过期时间
	Code       string `gorm:"column:code;type:varchar(8);default:''" json:"code"`                      // 最后发送成功的验证码
	Account    string `gorm:"column:account;type:varchar(100);default:''" json:"account"`              // 手机号或者邮箱
}

//get real primary key name
func (verificationCode *VerificationCode) GetKey() string {
	return "id"
}

//get primary key in model
func (verificationCode *VerificationCode) GetKeyProperty() uint64 {
	return verificationCode.Id
}

//set primary key
func (verificationCode *VerificationCode) SetKeyProperty(id uint64) {
	verificationCode.Id = id
}

//get real table name
func (verificationCode *VerificationCode) TableName() string {
	return "cmf_verification_code"
}
