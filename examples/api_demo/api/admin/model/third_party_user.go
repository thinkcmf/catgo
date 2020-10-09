package model

type ThirdPartyUser struct {
	Id            uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	UserId        uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:'0'" json:"user_id"`              // 本站用户id
	LastLoginTime uint   `gorm:"column:last_login_time;type:int(10) unsigned;default:'0'" json:"last_login_time"` // 最后登录时间
	ExpireTime    uint   `gorm:"column:expire_time;type:int(10) unsigned;default:'0'" json:"expire_time"`         // access_token过期时间
	CreateTime    uint   `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"`         // 绑定时间
	LoginTimes    uint   `gorm:"column:login_times;type:int(10) unsigned;default:'0'" json:"login_times"`         // 登录次数
	Status        uint8  `gorm:"column:status;type:tinyint(3) unsigned;default:'1'" json:"status"`                // 状态;1:正常;0:禁用
	Nickname      string `gorm:"column:nickname;type:varchar(50);default:''" json:"nickname"`                     // 用户昵称
	ThirdParty    string `gorm:"column:third_party;type:varchar(20);default:''" json:"third_party"`               // 第三方唯一码
	AppId         string `gorm:"column:app_id;type:varchar(64);default:''" json:"app_id"`                         // 第三方应用 id
	LastLoginIp   string `gorm:"column:last_login_ip;type:varchar(15);default:''" json:"last_login_ip"`           // 最后登录ip
	AccessToken   string `gorm:"column:access_token;type:varchar(512);default:''" json:"access_token"`            // 第三方授权码
	Openid        string `gorm:"column:openid;type:varchar(40);default:''" json:"openid"`                         // 第三方用户id
	UnionId       string `gorm:"column:union_id;type:varchar(64);default:''" json:"union_id"`                     // 第三方用户多个产品中的惟一 id,(如:微信平台)
	More          string `gorm:"column:more;type:text;default:null" json:"more"`                                  // 扩展信息
}

//get real primary key name
func (thirdPartyUser *ThirdPartyUser) GetKey() string {
	return "id"
}

//get primary key in model
func (thirdPartyUser *ThirdPartyUser) GetKeyProperty() uint64 {
	return thirdPartyUser.Id
}

//set primary key
func (thirdPartyUser *ThirdPartyUser) SetKeyProperty(id uint64) {
	thirdPartyUser.Id = id
}

//get real table name
func (thirdPartyUser *ThirdPartyUser) TableName() string {
	return "cmf_third_party_user"
}
