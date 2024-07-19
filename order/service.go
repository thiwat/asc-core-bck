package order

import (
	"asc-core/types"
	"errors"

	"asc-core/event"

	"go.mongodb.org/mongo-driver/bson"
)

func GetOrder(orderNo string, session types.Session) (Order, error) {
	return FindOne(bson.M{"order_no": orderNo, "user_id": session.UserId})
}

func ListByUser(page int64, pageSize int64, sort string, session types.Session) (ListOutput, error) {
	return List(bson.M{"user_id": session.UserId}, page, pageSize, sort)
}

func PlaceOrder(input PlaceOrderInput, session types.Session) (Order, error) {

	eventRec, err := event.FindByCode(input.Event)

	if err != nil {
		return Order{}, err
	}

	if input.Quantity > eventRec.AvailableSeat {
		return Order{}, errors.New("Exceed maximum quantity")
	}

	order := Order{
		UserId:   session.UserId,
		Event:    input.Event,
		Status:   string(PendingPayment),
		Quantity: input.Quantity,
	}

	res, err := Create(order)

	if err == nil {
		event.UpdateByCode(input.Event, bson.M{"available_seat": eventRec.AvailableSeat - input.Quantity})
	}

	return res, err
}

func UploadSlip(input UploadSlipInput, session types.Session) (Order, error) {

	order, err := FindOne(bson.M{
		"order_no": input.OrderNo,
		"user_id":  session.UserId,
	})

	if err != nil {
		return order, err
	}

	return UpdateOne(
		bson.M{"order_no": input.OrderNo, "user_id": session.UserId},
		bson.M{"slip_url": input.SlipUrl, "status": string(Paid)},
	)
}

func ApprovePayment(input ApprovePaymentInput) (Order, error) {

	order, err := FindOne(bson.M{
		"order_no": input.OrderNo,
	})

	if err != nil {
		return order, err
	}

	return UpdateOne(
		bson.M{"order_no": input.OrderNo},
		bson.M{"status": string(Completed)},
	)
}
