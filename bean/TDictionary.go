package bean

import "time"

type TDictionary struct {
	Id         int       `json:"id,omitemty"`
	Colname    string    `json:"colname,omitempty"`
	Value      int       `json:"value,omitemty"`
	Remark     string    `json:"remark,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
}
