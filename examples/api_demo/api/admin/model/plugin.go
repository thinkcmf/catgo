package model

type Plugin struct {
	Id          uint   `gorm:"column:id;type:int(11) unsigned" json:"id"`                               // 自增id
	Type        uint8  `gorm:"column:type;type:tinyint(3) unsigned;default:'1'" json:"type"`            // 插件类型;1:网站;8:微信
	HasAdmin    uint8  `gorm:"column:has_admin;type:tinyint(3) unsigned;default:'0'" json:"has_admin"`  // 是否有后台管理,0:没有;1:有
	Status      uint8  `gorm:"column:status;type:tinyint(3) unsigned;default:'1'" json:"status"`        // 状态;1:开启;0:禁用
	CreateTime  uint   `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"` // 插件安装时间
	Name        string `gorm:"column:name;type:varchar(50);default:''" json:"name"`                     // 插件标识名,英文字母(惟一)
	Title       string `gorm:"column:title;type:varchar(50);default:''" json:"title"`                   // 插件名称
	DemoUrl     string `gorm:"column:demo_url;type:varchar(50);default:''" json:"demo_url"`             // 演示地址，带协议
	Hooks       string `gorm:"column:hooks;type:varchar(255);default:''" json:"hooks"`                  // 实现的钩子;以“,”分隔
	Author      string `gorm:"column:author;type:varchar(20);default:''" json:"author"`                 // 插件作者
	AuthorUrl   string `gorm:"column:author_url;type:varchar(50);default:''" json:"author_url"`         // 作者网站链接
	Version     string `gorm:"column:version;type:varchar(20);default:''" json:"version"`               // 插件版本号
	Description string `gorm:"column:description;type:varchar(255)" json:"description"`                 // 插件描述
	Config      string `gorm:"column:config;type:text;default:null" json:"config"`                      // 插件配置
}

//get real primary key name
func (plugin *Plugin) GetKey() string {
	return "id"
}

//get primary key in model
func (plugin *Plugin) GetKeyProperty() uint {
	return plugin.Id
}

//set primary key
func (plugin *Plugin) SetKeyProperty(id uint) {
	plugin.Id = id
}

//get real table name
func (plugin *Plugin) TableName() string {
	return "cmf_plugin"
}
