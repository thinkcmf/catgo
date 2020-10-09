package model

type Hook struct {
	Id          uint   `gorm:"column:id;type:int(10) unsigned" json:"id"`
	Type        uint8  `gorm:"column:type;type:tinyint(3) unsigned;default:'0'" json:"type"`       // 钩子类型(1:系统钩子;2:应用钩子;3:模板钩子;4:后台模板钩子)
	Once        uint8  `gorm:"column:once;type:tinyint(3) unsigned;default:'0'" json:"once"`       // 是否只允许一个插件运行(0:多个;1:一个)
	Name        string `gorm:"column:name;type:varchar(50);default:''" json:"name"`                // 钩子名称
	Hook        string `gorm:"column:hook;type:varchar(50);default:''" json:"hook"`                // 钩子
	App         string `gorm:"column:app;type:varchar(15);default:''" json:"app"`                  // 应用名(只有应用钩子才用)
	Description string `gorm:"column:description;type:varchar(255);default:''" json:"description"` // 描述
}

//get real primary key name
func (hook *Hook) GetKey() string {
	return "id"
}

//get primary key in model
func (hook *Hook) GetKeyProperty() uint {
	return hook.Id
}

//set primary key
func (hook *Hook) SetKeyProperty(id uint) {
	hook.Id = id
}

//get real table name
func (hook *Hook) TableName() string {
	return "cmf_hook"
}
