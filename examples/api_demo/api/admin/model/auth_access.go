package model

type AuthAccess struct {
	Id       uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	RoleId   uint   `gorm:"column:role_id;type:int(10) unsigned" json:"role_id"`            // 角色
	RuleName string `gorm:"column:rule_name;type:varchar(100);default:''" json:"rule_name"` // 规则唯一英文标识,全小写
	Type     string `gorm:"column:type;type:varchar(30);default:''" json:"type"`            // 权限规则分类,请加应用前缀,如admin_
}

//get real primary key name
func (authAccess *AuthAccess) GetKey() string {
	return "id"
}

//get primary key in model
func (authAccess *AuthAccess) GetKeyProperty() uint64 {
	return authAccess.Id
}

//set primary key
func (authAccess *AuthAccess) SetKeyProperty(id uint64) {
	authAccess.Id = id
}

//get real table name
func (authAccess *AuthAccess) TableName() string {
	return "cmf_auth_access"
}
