package bean

import "time"

type TRelationship struct {
	Id               int       `json:"id,omitemty"`
	UserId1          int       `json:"user_id,omitemty"`
	UserId2          int       `json:"other_user_id,omitemty"`
	State            string    `json:"state,omitemty"`
	RelationshipType string    `json:"type,omitempty"`
	CreateTime       time.Time `json:"create_time,omitempty"`
	LastUpdateTime   time.Time `json:"last_update_time,omitempty"`
}

type RelationshipCoreInfo struct {
	UserId           int    `json:"user_id,omitemty"`
	State            string `json:"state,omitemty"`
	RelationshipType string `json:"type,omitempty"`
}
