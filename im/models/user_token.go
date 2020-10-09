package models

type UserToken struct {
	Id         int64
	UserId     int64
	ExpireTime int64
	CreateTime int64
	Token      string
	DeviceType string
}
