package ticket

import "time"

type Ticket struct {
	UserId    string    `json:"user_id" bson:"user_id"`
	Code      string    `json:"code" bson:"code"`
	Event     string    `json:"event" bson:"event"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
