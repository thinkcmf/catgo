package model

type UserAction struct {
	Id           int    `gorm:"column:id;type:int(11)" json:"id"`
	Score        int    `gorm:"column:score;type:int(11);default:'0'" json:"score"`                       // 更改积分，可以为负
	Coin         int    `gorm:"column:coin;type:int(11);default:'0'" json:"coin"`                         // 更改金币，可以为负
	RewardNumber int    `gorm:"column:reward_number;type:int(11);default:'0'" json:"reward_number"`       // 奖励次数
	CycleType    uint8  `gorm:"column:cycle_type;type:tinyint(1) unsigned;default:'0'" json:"cycle_type"` // 周期类型;0:不限;1:按天;2:按小时;3:永久
	CycleTime    uint   `gorm:"column:cycle_time;type:int(10) unsigned;default:'0'" json:"cycle_time"`    // 周期时间值
	Name         string `gorm:"column:name;type:varchar(50);default:''" json:"name"`                      // 用户操作名称
	Action       string `gorm:"column:action;type:varchar(50);default:''" json:"action"`                  // 用户操作名称
	App          string `gorm:"column:app;type:varchar(50);default:''" json:"app"`                        // 操作所在应用名或插件名等
	Url          string `gorm:"column:url;type:text;default:null" json:"url"`                             // 执行操作的url
}

//get real primary key name
func (userAction *UserAction) GetKey() string {
	return "id"
}

//get primary key in model
func (userAction *UserAction) GetKeyProperty() int {
	return userAction.Id
}

//set primary key
func (userAction *UserAction) SetKeyProperty(id int) {
	userAction.Id = id
}

//get real table name
func (userAction *UserAction) TableName() string {
	return "cmf_user_action"
}
