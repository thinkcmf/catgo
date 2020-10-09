package model

type UserToken struct {
	Id         int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	UserId     int64  `gorm:"column:user_id;type:bigint(20);default:'0'" json:"user_id"`               // 用户id
	ExpireTime uint   `gorm:"column:expire_time;type:int(10) unsigned;default:'0'" json:"expire_time"` //  过期时间
	CreateTime uint   `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"` // 创建时间
	Token      string `gorm:"column:token;type:varchar(64);default:''" json:"token"`                   // token
	DeviceType string `gorm:"column:device_type;type:varchar(10);default:''" json:"device_type"`       // 设备类型;mobile,android,iphone,ipad,web,pc,mac,wxapp
}

//get real primary key name
func (userToken *UserToken) GetKey() string {
	return "id"
}

//get primary key in model
func (userToken *UserToken) GetKeyProperty() int64 {
	return userToken.Id
}

//set primary key
func (userToken *UserToken) SetKeyProperty(id int64) {
	userToken.Id = id
}

//get real table name
func (userToken *UserToken) TableName() string {
	return "cmf_user_token"
}
