package model

type ThemeFile struct {
	Id          int     `gorm:"column:id;type:int(11)" json:"id"`
	IsPublic    int8    `gorm:"column:is_public;type:tinyint(4);default:'0'" json:"is_public"`      // 是否公共的模板文件
	ListOrder   float32 `gorm:"column:list_order;type:float;default:'10000'" json:"list_order"`     // 排序
	Theme       string  `gorm:"column:theme;type:varchar(20);default:''" json:"theme"`              // 模板名称
	Name        string  `gorm:"column:name;type:varchar(20);default:''" json:"name"`                // 模板文件名
	Action      string  `gorm:"column:action;type:varchar(50);default:''" json:"action"`            // 操作
	File        string  `gorm:"column:file;type:varchar(50);default:''" json:"file"`                // 模板文件，相对于模板根目录，如Portal/index.html
	Description string  `gorm:"column:description;type:varchar(100);default:''" json:"description"` // 模板文件描述
	More        string  `gorm:"column:more;type:text;default:null" json:"more"`                     // 模板更多配置,用户自己后台设置的
	ConfigMore  string  `gorm:"column:config_more;type:text;default:null" json:"config_more"`       // 模板更多配置,来源模板的配置文件
	DraftMore   string  `gorm:"column:draft_more;type:text;default:null" json:"draft_more"`         // 模板更多配置,用户临时保存的配置
}

//get real primary key name
func (themeFile *ThemeFile) GetKey() string {
	return "id"
}

//get primary key in model
func (themeFile *ThemeFile) GetKeyProperty() int {
	return themeFile.Id
}

//set primary key
func (themeFile *ThemeFile) SetKeyProperty(id int) {
	themeFile.Id = id
}

//get real table name
func (themeFile *ThemeFile) TableName() string {
	return "cmf_theme_file"
}
