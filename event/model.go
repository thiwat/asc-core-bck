package event

import "time"

type Event struct {
	Name           string    `json:"name" bson:"name"`
	Code           string    `json:"code" bson:"code"`
	NumberOfSeat   int32     `json:"number_of_seat" bson:"number_of_seat"`
	AvailableSeat  int32     `json:"available_seat" bson:"available_seat"`
	Price          int64     `json:"price" bson:"price"`
	SaleStartDate  time.Time `json:"sale_start_date" bson:"sale_start_date,omitempty"`
	SaleEndDate    time.Time `json:"sale_end_date" bson:"sale_end_date,omitempty"`
	EventStartTime time.Time `json:"event_start_time" bson:"event_start_time,omitempty"`
	EventEndTime   time.Time `json:"event_end_time" bson:"event_end_time,omitempty"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
