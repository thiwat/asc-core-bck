package order

import "time"

type Order struct {
	OrderNo   string    `json:"order_no" bson:"order_no"`
	UserId    string    `json:"user_id" bson:"user_id"`
	Event     string    `json:"event" bson:"event"`
	Status    string    `json:"status" bson:"status"`
	Quantity  int32     `json:"quantity" bson:"quantity"`
	SlipUrl   string    `json:"slip_url" bson:"slip_url"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
