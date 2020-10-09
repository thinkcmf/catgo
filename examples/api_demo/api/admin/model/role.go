package model

type Role struct {
	Id         uint    `gorm:"column:id;type:int(10) unsigned" json:"id"`
	ParentId   uint    `gorm:"column:parent_id;type:int(10) unsigned;default:'0'" json:"parent_id"`     // 父角色ID
	Status     uint8   `gorm:"column:status;type:tinyint(3) unsigned;default:'0'" json:"status"`        // 状态;0:禁用;1:正常
	CreateTime uint    `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"` // 创建时间
	UpdateTime uint    `gorm:"column:update_time;type:int(10) unsigned;default:'0'" json:"update_time"` // 更新时间
	ListOrder  float32 `gorm:"column:list_order;type:float;default:'0'" json:"list_order"`              // 排序
	Name       string  `gorm:"column:name;type:varchar(20);default:''" json:"name"`                     // 角色名称
	Remark     string  `gorm:"column:remark;type:varchar(255);default:''" json:"remark"`                // 备注
}

//get real primary key name
func (role *Role) GetKey() string {
	return "id"
}

//get primary key in model
func (role *Role) GetKeyProperty() uint {
	return role.Id
}

//set primary key
func (role *Role) SetKeyProperty(id uint) {
	role.Id = id
}

//get real table name
func (role *Role) TableName() string {
	return "cmf_role"
}
