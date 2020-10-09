package model

type AuthRule struct {
	Id        uint   `gorm:"column:id;type:int(10) unsigned" json:"id"`                        // 规则id,自增主键
	Status    uint8  `gorm:"column:status;type:tinyint(3) unsigned;default:'1'" json:"status"` // 是否有效(0:无效,1:有效)
	App       string `gorm:"column:app;type:varchar(40);default:''" json:"app"`                // 规则所属app
	Type      string `gorm:"column:type;type:varchar(30);default:''" json:"type"`              // 权限规则分类，请加应用前缀,如admin_
	Name      string `gorm:"column:name;type:varchar(100);default:''" json:"name"`             // 规则唯一英文标识,全小写
	Param     string `gorm:"column:param;type:varchar(100);default:''" json:"param"`           // 额外url参数
	Title     string `gorm:"column:title;type:varchar(20);default:''" json:"title"`            // 规则描述
	Condition string `gorm:"column:condition;type:varchar(200);default:''" json:"condition"`   // 规则附加条件
}

//get real primary key name
func (authRule *AuthRule) GetKey() string {
	return "id"
}

//get primary key in model
func (authRule *AuthRule) GetKeyProperty() uint {
	return authRule.Id
}

//set primary key
func (authRule *AuthRule) SetKeyProperty(id uint) {
	authRule.Id = id
}

//get real table name
func (authRule *AuthRule) TableName() string {
	return "cmf_auth_rule"
}
