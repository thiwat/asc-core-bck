package event

import "time"

type Event struct {
	Name           string    `json:"name" bson:"name,omitempty"`
	Code           string    `json:"code" bson:"code,omitempty"`
	Location       string    `json:"location" bson:"location,omitempty"`
	Detail         string    `json:"detail" bson:"detail,omitempty"`
	Cover          string    `json:"cover" bson:"cover,omitempty"`
	NumberOfSeat   int32     `json:"number_of_seat" bson:"number_of_seat,omitempty"`
	AvailableSeat  int32     `json:"available_seat" bson:"available_seat,omitempty"`
	Price          int64     `json:"price" bson:"price,omitempty"`
	SaleStartDate  time.Time `json:"sale_start_date" bson:"sale_start_date,omitempty"`
	SaleEndDate    time.Time `json:"sale_end_date" bson:"sale_end_date,omitempty"`
	EventStartDate time.Time `json:"event_start_date" bson:"event_start_date,omitempty"`
	EventEndDate   time.Time `json:"event_end_date" bson:"event_end_date,omitempty"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
