package model

type Option struct {
	Id          uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	Autoload    uint8  `gorm:"column:autoload;type:tinyint(3) unsigned;default:'1'" json:"autoload"` // 是否自动加载;1:自动加载;0:不自动加载
	OptionName  string `gorm:"column:option_name;type:varchar(64);default:''" json:"option_name"`    // 配置名
	OptionValue string `gorm:"column:option_value;type:longtext;default:null" json:"option_value"`   // 配置值
}

//get real primary key name
func (option *Option) GetKey() string {
	return "id"
}

//get primary key in model
func (option *Option) GetKeyProperty() uint64 {
	return option.Id
}

//set primary key
func (option *Option) SetKeyProperty(id uint64) {
	option.Id = id
}

//get real table name
func (option *Option) TableName() string {
	return "cmf_option"
}
