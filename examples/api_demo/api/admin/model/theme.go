package model

type Theme struct {
	Id          int    `gorm:"column:id;type:int(11)" json:"id"`
	CreateTime  uint   `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"`    // 安装时间
	UpdateTime  uint   `gorm:"column:update_time;type:int(10) unsigned;default:'0'" json:"update_time"`    // 最后升级时间
	Status      uint8  `gorm:"column:status;type:tinyint(3) unsigned;default:'0'" json:"status"`           // 模板状态,1:正在使用;0:未使用
	IsCompiled  uint8  `gorm:"column:is_compiled;type:tinyint(3) unsigned;default:'0'" json:"is_compiled"` // 是否为已编译模板
	Theme       string `gorm:"column:theme;type:varchar(20);default:''" json:"theme"`                      // 主题目录名，用于主题的维一标识
	Name        string `gorm:"column:name;type:varchar(20);default:''" json:"name"`                        // 主题名称
	Version     string `gorm:"column:version;type:varchar(20);default:''" json:"version"`                  // 主题版本号
	DemoUrl     string `gorm:"column:demo_url;type:varchar(50);default:''" json:"demo_url"`                // 演示地址，带协议
	Thumbnail   string `gorm:"column:thumbnail;type:varchar(100);default:''" json:"thumbnail"`             // 缩略图
	Author      string `gorm:"column:author;type:varchar(20);default:''" json:"author"`                    // 主题作者
	AuthorUrl   string `gorm:"column:author_url;type:varchar(50);default:''" json:"author_url"`            // 作者网站链接
	Lang        string `gorm:"column:lang;type:varchar(10);default:''" json:"lang"`                        // 支持语言
	Keywords    string `gorm:"column:keywords;type:varchar(50);default:''" json:"keywords"`                // 主题关键字
	Description string `gorm:"column:description;type:varchar(100);default:''" json:"description"`         // 主题描述
}

//get real primary key name
func (theme *Theme) GetKey() string {
	return "id"
}

//get primary key in model
func (theme *Theme) GetKeyProperty() int {
	return theme.Id
}

//set primary key
func (theme *Theme) SetKeyProperty(id int) {
	theme.Id = id
}

//get real table name
func (theme *Theme) TableName() string {
	return "cmf_theme"
}
