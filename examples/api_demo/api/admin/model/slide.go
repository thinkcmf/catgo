package model

type Slide struct {
	Id         int    `gorm:"column:id;type:int(11)" json:"id"`
	Status     uint8  `gorm:"column:status;type:tinyint(3) unsigned;default:'1'" json:"status"`        // 状态,1:显示,0不显示
	DeleteTime uint   `gorm:"column:delete_time;type:int(10) unsigned;default:'0'" json:"delete_time"` // 删除时间
	Name       string `gorm:"column:name;type:varchar(50);default:''" json:"name"`                     // 幻灯片分类
	Remark     string `gorm:"column:remark;type:varchar(255);default:''" json:"remark"`                // 分类备注
}

//get real primary key name
func (slide *Slide) GetKey() string {
	return "id"
}

//get primary key in model
func (slide *Slide) GetKeyProperty() int {
	return slide.Id
}

//set primary key
func (slide *Slide) SetKeyProperty(id int) {
	slide.Id = id
}

//get real table name
func (slide *Slide) TableName() string {
	return "cmf_slide"
}
