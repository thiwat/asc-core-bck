package ticket

import (
	"errors"
	"asc-core/types"

	"go.mongodb.org/mongo-driver/bson"
)

func GetTicket(code string) (Ticket, error) {
	return FindOne(bson.M{"code": code})
}

func PurchaseTicket(input PurchaseTicketInput, session types.Session) (Ticket, error) {

	exists, _ := FindOne(bson.M{"user_id": session.UserId, "event": input.Event})

	if exists != (Ticket{}) {
		return Ticket{}, errors.New("already_purchase")
	}

	ticket := Ticket{
		UserId: session.UserId,
		Event: input.Event,
	}

	return Create(ticket)
}

func FindTicketByUser(event string, session types.Session) (Ticket, error) {
	return FindOne(bson.M{"event": event, "user_id": session.UserId})
}

func ListTicketByUser(page int64, pageSize int64, sort string, session types.Session) (types.ListOutput[Ticket], error) {
	filter := bson.M{"user_id": session.UserId}
	return List(filter, page, pageSize, sort)
}