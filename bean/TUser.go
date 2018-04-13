package bean

import "time"

type TUser struct {
	Id             int       `json:"id,omitemty"`
	Name           string    `json:"name,omitempty"`
	UserType       string    `json:"type,omitempty"`
	CreateTime     time.Time `json:"create_time,omitempty"`
	LastUpdateTime time.Time `json:"last_update_time,omitempty"`
}

type UserCoreInfo struct {
	Id       int    `json:"id,omitemty"`
	Name     string `json:"name,omitempty"`
	UserType string `json:"type,omitempty"`
}
