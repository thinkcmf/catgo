package models

type UserToken struct {
	Id         int64
	UserId     string
	ExpireTime string
	CreateTime int64
	Token      string
	DeviceType string
}
