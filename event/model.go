package event

import "time"

type Event struct {
	Name      string    `json:"name" bson:"name"`
	Code      string    `json:"code" bson:"code"`
	Seat      int32     `json:"seat" bson:"seat"`
	Price     int64     `json:"price" bson:"price"`
	StartDate time.Time `json:"start_date" bson:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date" bson:"end_date,omitempty"`
	ShowDay   time.Time `json:"show_day" bson:"show_day,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
