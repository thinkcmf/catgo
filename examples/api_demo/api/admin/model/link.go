package model

type Link struct {
	Id          uint64  `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	Status      uint8   `gorm:"column:status;type:tinyint(3) unsigned;default:'1'" json:"status"`   // 状态;1:显示;0:不显示
	Rating      int     `gorm:"column:rating;type:int(11);default:'0'" json:"rating"`               // 友情链接评级
	ListOrder   float32 `gorm:"column:list_order;type:float;default:'10000'" json:"list_order"`     // 排序
	Description string  `gorm:"column:description;type:varchar(255);default:''" json:"description"` // 友情链接描述
	Url         string  `gorm:"column:url;type:varchar(255);default:''" json:"url"`                 // 友情链接地址
	Name        string  `gorm:"column:name;type:varchar(30);default:''" json:"name"`                // 友情链接名称
	Image       string  `gorm:"column:image;type:varchar(100);default:''" json:"image"`             // 友情链接图标
	Target      string  `gorm:"column:target;type:varchar(10);default:''" json:"target"`            // 友情链接打开方式
	Rel         string  `gorm:"column:rel;type:varchar(50);default:''" json:"rel"`                  // 链接与网站的关系
}

//get real primary key name
func (link *Link) GetKey() string {
	return "id"
}

//get primary key in model
func (link *Link) GetKeyProperty() uint64 {
	return link.Id
}

//set primary key
func (link *Link) SetKeyProperty(id uint64) {
	link.Id = id
}

//get real table name
func (link *Link) TableName() string {
	return "cmf_link"
}
