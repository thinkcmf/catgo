package model

type UserFavorite struct {
	Id          uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	UserId      uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:'0'" json:"user_id"`      // 用户 id
	Title       string `gorm:"column:title;type:varchar(100);default:''" json:"title"`                  // 收藏内容的标题
	Thumbnail   string `gorm:"column:thumbnail;type:varchar(100);default:''" json:"thumbnail"`          // 缩略图
	Url         string `gorm:"column:url;type:varchar(255);default:null" json:"url"`                    // 收藏内容的原文地址，JSON格式
	Description string `gorm:"column:description;type:text;default:null" json:"description"`            // 收藏内容的描述
	Table       string `gorm:"column:table_name;type:varchar(64);default:''" json:"table_name"`         // 收藏实体以前所在表,不带前缀
	ObjectId    uint   `gorm:"column:object_id;type:int(10) unsigned;default:'0'" json:"object_id"`     // 收藏内容原来的主键id
	CreateTime  uint   `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"` // 收藏时间
}

//get real primary key name
func (userFavorite *UserFavorite) GetKey() string {
	return "id"
}

//get primary key in model
func (userFavorite *UserFavorite) GetKeyProperty() uint64 {
	return userFavorite.Id
}

//set primary key
func (userFavorite *UserFavorite) SetKeyProperty(id uint64) {
	userFavorite.Id = id
}

//get real table name
func (userFavorite *UserFavorite) TableName() string {
	return "cmf_user_favorite"
}
