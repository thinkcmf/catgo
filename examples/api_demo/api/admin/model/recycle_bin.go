package model

type RecycleBin struct {
	Id         uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	ObjectId   int    `gorm:"column:object_id;type:int(11);default:'0'" json:"object_id"`              // 删除内容 id
	CreateTime uint   `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"` // 创建时间
	Table      string `gorm:"column:table_name;type:varchar(60);default:''" json:"table_name"`         // 删除内容所在表名
	Name       string `gorm:"column:name;type:varchar(255);default:''" json:"name"`                    // 删除内容名称
	UserId     uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:'0'" json:"user_id"`      // 用户id
}

//get real primary key name
func (recycleBin *RecycleBin) GetKey() string {
	return "id"
}

//get primary key in model
func (recycleBin *RecycleBin) GetKeyProperty() uint64 {
	return recycleBin.Id
}

//set primary key
func (recycleBin *RecycleBin) SetKeyProperty(id uint64) {
	recycleBin.Id = id
}

//get real table name
func (recycleBin *RecycleBin) TableName() string {
	return "cmf_recycle_bin"
}
