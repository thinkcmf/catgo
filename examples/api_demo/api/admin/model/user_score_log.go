package model

type UserScoreLog struct {
	Id         int    `gorm:"column:id;type:int(11)" json:"id"`
	UserId     int64  `gorm:"column:user_id;type:bigint(20);default:'0'" json:"user_id"`      // 用户 id
	CreateTime int    `gorm:"column:create_time;type:int(11);default:'0'" json:"create_time"` // 创建时间
	Action     string `gorm:"column:action;type:varchar(50);default:''" json:"action"`        // 用户操作名称
	Score      int    `gorm:"column:score;type:int(11);default:'0'" json:"score"`             // 更改积分，可以为负
	Coin       int    `gorm:"column:coin;type:int(11);default:'0'" json:"coin"`               // 更改金币，可以为负
}

//get real primary key name
func (userScoreLog *UserScoreLog) GetKey() string {
	return "id"
}

//get primary key in model
func (userScoreLog *UserScoreLog) GetKeyProperty() int {
	return userScoreLog.Id
}

//set primary key
func (userScoreLog *UserScoreLog) SetKeyProperty(id int) {
	userScoreLog.Id = id
}

//get real table name
func (userScoreLog *UserScoreLog) TableName() string {
	return "cmf_user_score_log"
}
