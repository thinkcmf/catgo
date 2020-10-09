package model

type RoleUser struct {
	Id     uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	RoleId uint   `gorm:"column:role_id;type:int(10) unsigned;default:'0'" json:"role_id"`    // 角色 id
	UserId uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:'0'" json:"user_id"` // 用户id
}

//get real primary key name
func (roleUser *RoleUser) GetKey() string {
	return "id"
}

//get primary key in model
func (roleUser *RoleUser) GetKeyProperty() uint64 {
	return roleUser.Id
}

//set primary key
func (roleUser *RoleUser) SetKeyProperty(id uint64) {
	roleUser.Id = id
}

//get real table name
func (roleUser *RoleUser) TableName() string {
	return "cmf_role_user"
}
