package order

import (
	"asc-core/types"
	"errors"

	"asc-core/event"

	"go.mongodb.org/mongo-driver/bson"
)

func GetOrder(orderNo string, session types.Session) (Order, error) {
	return findOne(bson.M{"order_no": orderNo, "user_id": session.UserId})
}

func ListByUser(page int64, pageSize int64, sort string, session types.Session) (ListOutput, error) {
	return list(bson.M{"user_id": session.UserId}, page, pageSize, sort)
}

func PlaceOrder(input PlaceOrderInput, session types.Session) (Order, error) {

	eventRec, err := event.GetEvent(input.Event)

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

	res, err := create(order)

	if err == nil {
		event.UpdateByCode(input.Event, event.Event{AvailableSeat: eventRec.AvailableSeat - input.Quantity})
	}

	return res, err
}

func UploadSlip(input UploadSlipInput, session types.Session) (Order, error) {

	order, err := findOne(bson.M{
		"order_no": input.OrderNo,
		"user_id":  session.UserId,
	})

	if err != nil {
		return order, err
	}

	return updateOne(
		bson.M{"order_no": input.OrderNo, "user_id": session.UserId},
		Order{SlipUrl: input.SlipUrl, Status: string(Paid)},
	)
}

func ApprovePayment(input ApprovePaymentInput) (Order, error) {

	order, err := findOne(bson.M{
		"order_no": input.OrderNo,
	})

	if err != nil {
		return order, err
	}

	return updateOne(
		bson.M{"order_no": input.OrderNo},
		Order{Status: string(Completed)},
	)
}
