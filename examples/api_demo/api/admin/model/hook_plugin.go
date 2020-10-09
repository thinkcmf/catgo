package model

type HookPlugin struct {
	Id        uint    `gorm:"column:id;type:int(10) unsigned" json:"id"`
	ListOrder float32 `gorm:"column:list_order;type:float;default:'10000'" json:"list_order"` // 排序
	Status    int8    `gorm:"column:status;type:tinyint(4);default:'1'" json:"status"`        // 状态(0:禁用,1:启用)
	Hook      string  `gorm:"column:hook;type:varchar(50);default:''" json:"hook"`            // 钩子名
	Plugin    string  `gorm:"column:plugin;type:varchar(50);default:''" json:"plugin"`        // 插件
}

//get real primary key name
func (hookPlugin *HookPlugin) GetKey() string {
	return "id"
}

//get primary key in model
func (hookPlugin *HookPlugin) GetKeyProperty() uint {
	return hookPlugin.Id
}

//set primary key
func (hookPlugin *HookPlugin) SetKeyProperty(id uint) {
	hookPlugin.Id = id
}

//get real table name
func (hookPlugin *HookPlugin) TableName() string {
	return "cmf_hook_plugin"
}
