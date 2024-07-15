package user

import "time"

type User struct {
	UserId       string    `json:"user_id" bson:"user_id"`
	LineId       string    `json:"line_id" bson:"line_id"`
	Name         string    `json:"name" bson:"name"`
	ProfileImage string    `json:"profile_image" bson:"profile_image"`
	RegisteredAt time.Time `json:"registered_at" bson:"registered_at,omitempty"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
