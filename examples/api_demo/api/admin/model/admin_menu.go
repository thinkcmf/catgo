package model

type AdminMenu struct {
	Id         uint    `gorm:"column:id;type:int(10) unsigned" json:"id"`
	ParentId   uint    `gorm:"column:parent_id;type:int(10) unsigned;default:'0'" json:"parent_id"` // 父菜单id
	Type       uint8   `gorm:"column:type;type:tinyint(3) unsigned;default:'1'" json:"type"`        // 菜单类型;1:有界面可访问菜单,2:无界面可访问菜单,0:只作为菜单
	Status     uint8   `gorm:"column:status;type:tinyint(3) unsigned;default:'0'" json:"status"`    // 状态;1:显示,0:不显示
	ListOrder  float32 `gorm:"column:list_order;type:float;default:'10000'" json:"list_order"`      // 排序
	App        string  `gorm:"column:app;type:varchar(40);default:''" json:"app"`                   // 应用名
	Controller string  `gorm:"column:controller;type:varchar(30);default:''" json:"controller"`     // 控制器名
	Action     string  `gorm:"column:action;type:varchar(30);default:''" json:"action"`             // 操作名称
	Param      string  `gorm:"column:param;type:varchar(50);default:''" json:"param"`               // 额外参数
	Name       string  `gorm:"column:name;type:varchar(30);default:''" json:"name"`                 // 菜单名称
	Icon       string  `gorm:"column:icon;type:varchar(20);default:''" json:"icon"`                 // 菜单图标
	Remark     string  `gorm:"column:remark;type:varchar(255);default:''" json:"remark"`            // 备注
}

//get real primary key name
func (adminMenu *AdminMenu) GetKey() string {
	return "id"
}

//get primary key in model
func (adminMenu *AdminMenu) GetKeyProperty() uint {
	return adminMenu.Id
}

//set primary key
func (adminMenu *AdminMenu) SetKeyProperty(id uint) {
	adminMenu.Id = id
}

//get real table name
func (adminMenu *AdminMenu) TableName() string {
	return "cmf_admin_menu"
}
