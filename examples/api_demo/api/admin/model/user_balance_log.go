package model

type UserBalanceLog struct {
	Id          uint    `gorm:"column:id;type:int(10) unsigned" json:"id"`
	UserId      int64   `gorm:"column:user_id;type:bigint(20);default:'0'" json:"user_id"`               // 用户 id
	CreateTime  uint    `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"` // 创建时间
	Change      float64 `gorm:"column:change;type:decimal(10,2);default:'0.00'" json:"change"`           // 更改余额
	Balance     float64 `gorm:"column:balance;type:decimal(10,2);default:'0.00'" json:"balance"`         // 更改后余额
	Description string  `gorm:"column:description;type:varchar(255);default:''" json:"description"`      // 描述
	Remark      string  `gorm:"column:remark;type:varchar(255);default:''" json:"remark"`                // 备注
}

//get real primary key name
func (userBalanceLog *UserBalanceLog) GetKey() string {
	return "id"
}

//get primary key in model
func (userBalanceLog *UserBalanceLog) GetKeyProperty() uint {
	return userBalanceLog.Id
}

//set primary key
func (userBalanceLog *UserBalanceLog) SetKeyProperty(id uint) {
	userBalanceLog.Id = id
}

//get real table name
func (userBalanceLog *UserBalanceLog) TableName() string {
	return "cmf_user_balance_log"
}
