package model

type UserLike struct {
	Id          uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	UserId      uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:'0'" json:"user_id"`      // 用户 id
	ObjectId    uint   `gorm:"column:object_id;type:int(10) unsigned;default:'0'" json:"object_id"`     // 内容原来的主键id
	CreateTime  uint   `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"` // 创建时间
	TableName   string `gorm:"column:table_name;type:varchar(64);default:''" json:"table_name"`         // 内容以前所在表,不带前缀
	Url         string `gorm:"column:url;type:varchar(255);default:''" json:"url"`                      // 内容的原文地址，不带域名
	Title       string `gorm:"column:title;type:varchar(100);default:''" json:"title"`                  // 内容的标题
	Thumbnail   string `gorm:"column:thumbnail;type:varchar(100);default:''" json:"thumbnail"`          // 缩略图
	Description string `gorm:"column:description;type:text;default:null" json:"description"`            // 内容的描述
}

//get real primary key name
func (userLike *UserLike) GetKey() string {
	return "id"
}

//get primary key in model
func (userLike *UserLike) GetKeyProperty() uint64 {
	return userLike.Id
}

//set primary key
func (userLike *UserLike) SetKeyProperty(id uint64) {
	userLike.Id = id
}

//get real table name
func (userLike *UserLike) TableName() string {
	return "cmf_user_like"
}
