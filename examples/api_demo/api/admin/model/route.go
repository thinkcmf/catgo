package model

type Route struct {
	Id        int     `gorm:"column:id;type:int(11)" json:"id"`                               // 路由id
	ListOrder float32 `gorm:"column:list_order;type:float;default:'10000'" json:"list_order"` // 排序
	Status    int8    `gorm:"column:status;type:tinyint(2);default:'1'" json:"status"`        // 状态;1:启用,0:不启用
	Type      int8    `gorm:"column:type;type:tinyint(4);default:'1'" json:"type"`            // URL规则类型;1:用户自定义;2:别名添加
	FullUrl   string  `gorm:"column:full_url;type:varchar(255);default:''" json:"full_url"`   // 完整url
	Url       string  `gorm:"column:url;type:varchar(255);default:''" json:"url"`             // 实际显示的url
}

//get real primary key name
func (route *Route) GetKey() string {
	return "id"
}

//get primary key in model
func (route *Route) GetKeyProperty() int {
	return route.Id
}

//set primary key
func (route *Route) SetKeyProperty(id int) {
	route.Id = id
}

//get real table name
func (route *Route) TableName() string {
	return "cmf_route"
}
