package user

import "time"

type User struct {
	UserId       string    `json:"user_id" bson:"user_id,omitempty"`
	LineId       string    `json:"line_id" bson:"line_id,omitempty"`
	Name         string    `json:"name" bson:"name,omitempty"`
	ProfileImage string    `json:"profile_image" bson:"profile_image,omitempty"`
	RegisteredAt time.Time `json:"registered_at" bson:"registered_at,omitempty"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
