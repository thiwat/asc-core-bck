package order

import "time"

type Order struct {
	OrderNo   string    `json:"order_no" bson:"order_no,omitempty"`
	UserId    string    `json:"user_id" bson:"user_id,omitempty"`
	Event     string    `json:"event" bson:"event,omitempty"`
	Status    string    `json:"status" bson:"status,omitempty"`
	Quantity  int32     `json:"quantity" bson:"quantity,omitempty"`
	SlipUrl   string    `json:"slip_url" bson:"slip_url,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
