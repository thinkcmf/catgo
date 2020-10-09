package model

type Asset struct {
	Id            uint64 `gorm:"column:id;type:bigint(20) unsigned" json:"id"`
	UserId        uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:'0'" json:"user_id"`            // 用户id
	FileSize      uint64 `gorm:"column:file_size;type:bigint(20) unsigned;default:'0'" json:"file_size"`        // 文件大小,单位B
	CreateTime    uint   `gorm:"column:create_time;type:int(10) unsigned;default:'0'" json:"create_time"`       // 上传时间
	Status        uint8  `gorm:"column:status;type:tinyint(3) unsigned;default:'1'" json:"status"`              // 状态;1:可用,0:不可用
	DownloadTimes uint   `gorm:"column:download_times;type:int(10) unsigned;default:'0'" json:"download_times"` // 下载次数
	FileKey       string `gorm:"column:file_key;type:varchar(64);default:''" json:"file_key"`                   // 文件惟一码
	Filename      string `gorm:"column:filename;type:varchar(100);default:''" json:"filename"`                  // 文件名
	FilePath      string `gorm:"column:file_path;type:varchar(100);default:''" json:"file_path"`                // 文件路径,相对于upload目录,可以为url
	FileMd5       string `gorm:"column:file_md5;type:varchar(32);default:''" json:"file_md5"`                   // 文件md5值
	FileSha1      string `gorm:"column:file_sha1;type:varchar(40);default:''" json:"file_sha1"`
	Suffix        string `gorm:"column:suffix;type:varchar(10);default:''" json:"suffix"` // 文件后缀名,不包括点
	More          string `gorm:"column:more;type:text;default:null" json:"more"`          // 其它详细信息,JSON格式
}

//get real primary key name
func (asset *Asset) GetKey() string {
	return "id"
}

//get primary key in model
func (asset *Asset) GetKeyProperty() uint64 {
	return asset.Id
}

//set primary key
func (asset *Asset) SetKeyProperty(id uint64) {
	asset.Id = id
}

//get real table name
func (asset *Asset) TableName() string {
	return "cmf_asset"
}
