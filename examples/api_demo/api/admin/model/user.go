package model

type User struct {
	Id                uint64  `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	UserType          uint8   `gorm:"column:user_type;type:tinyint(3) unsigned;default:1" json:"user_type"`            // 用户类型;1:admin;2:会员
	Sex               int8    `gorm:"column:sex;type:tinyint(2);default:0" json:"sex"`                                 // 性别;0:保密,1:男,2:女
	Birthday          int     `gorm:"column:birthday;type:int(11);default:0" json:"birthday"`                          // 生日
	LastLoginTime     int     `gorm:"column:last_login_time;type:int(11);default:0" json:"last_login_time"`            // 最后登录时间
	Score             int     `gorm:"column:score;type:int(11);default:0" json:"score"`                                // 用户积分
	Coin              uint    `gorm:"column:coin;type:int(10) unsigned;default:0" json:"coin"`                         // 金币
	Balance           float64 `gorm:"column:balance;type:decimal(10,2);default:0.00" json:"balance"`                   // 余额
	CreateTime        int     `gorm:"column:create_time;type:int(11);default:0" json:"create_time"`                    // 注册时间
	UserStatus        uint8   `gorm:"column:user_status;type:tinyint(3) unsigned;default:1" json:"user_status"`        // 用户状态;0:禁用,1:正常,2:未验证
	UserLogin         string  `gorm:"column:user_login;type:varchar(60);default:''" json:"user_login"`                   // 用户名
	UserPass          string  `gorm:"column:user_pass;type:varchar(64);default:''" json:"user_pass"`                     // 登录密码;cmf_password加密
	UserNickname      string  `gorm:"column:user_nickname;type:varchar(50);default:''" json:"user_nickname"`             // 用户昵称
	UserEmail         string  `gorm:"column:user_email;type:varchar(100);default:''" json:"user_email"`                  // 用户登录邮箱
	UserUrl           string  `gorm:"column:user_url;type:varchar(100);default:''" json:"user_url"`                      // 用户个人网址
	Avatar            string  `gorm:"column:avatar;type:varchar(255);default:''" json:"avatar"`                          // 用户头像
	Signature         string  `gorm:"column:signature;type:varchar(255);default:''" json:"signature"`                    // 个性签名
	LastLoginIp       string  `gorm:"column:last_login_ip;type:varchar(15);default:''" json:"last_login_ip"`             // 最后登录ip
	UserActivationKey string  `gorm:"column:user_activation_key;type:varchar(60);default:''" json:"user_activation_key"` // 激活码
	Mobile            string  `gorm:"column:mobile;type:varchar(20);default:''" json:"mobile"`                           // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	More              string  `gorm:"column:more;type:text;default:null" json:"more"`                                    // 扩展属性
}

//get real primary key name
func (user *User) GetKey() string {
	return "id"
}

//get primary key in model
func (user *User) GetKeyProperty() uint64 {
	return user.Id
}

//set primary key
func (user *User) SetKeyProperty(id uint64) {
	user.Id = id
}

//get real table name
func (user *User) TableName() string {
	return "cmf_user"
}
