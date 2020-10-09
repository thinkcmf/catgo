package model

type SlideItem struct {
	Id          uint    `gorm:"column:id;type:int(10) unsigned" json:"id"`
	SlideId     int     `gorm:"column:slide_id;type:int(11);default:'0'" json:"slide_id"`         // 幻灯片id
	Status      uint8   `gorm:"column:status;type:tinyint(3) unsigned;default:'1'" json:"status"` // 状态,1:显示;0:隐藏
	ListOrder   float32 `gorm:"column:list_order;type:float;default:'10000'" json:"list_order"`   // 排序
	Title       string  `gorm:"column:title;type:varchar(50);default:''" json:"title"`            // 幻灯片名称
	Image       string  `gorm:"column:image;type:varchar(255);default:''" json:"image"`           // 幻灯片图片
	Url         string  `gorm:"column:url;type:varchar(255);default:''" json:"url"`               // 幻灯片链接
	Target      string  `gorm:"column:target;type:varchar(10);default:''" json:"target"`          // 友情链接打开方式
	Description string  `gorm:"column:description;type:varchar(255)" json:"description"`          // 幻灯片描述
	Content     string  `gorm:"column:content;type:text;default:null" json:"content"`             // 幻灯片内容
	More        string  `gorm:"column:more;type:text;default:null" json:"more"`                   // 扩展信息
}

//get real primary key name
func (slideItem *SlideItem) GetKey() string {
	return "id"
}

//get primary key in model
func (slideItem *SlideItem) GetKeyProperty() uint {
	return slideItem.Id
}

//set primary key
func (slideItem *SlideItem) SetKeyProperty(id uint) {
	slideItem.Id = id
}

//get real table name
func (slideItem *SlideItem) TableName() string {
	return "cmf_slide_item"
}
