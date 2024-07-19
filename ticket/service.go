package ticket

import (
	"asc-core/types"

	"go.mongodb.org/mongo-driver/bson"
)

func GetTicket(code string) (Ticket, error) {
	return FindOne(bson.M{"code": code})
}

func FindTicketByUser(event string, session types.Session) (Ticket, error) {
	return FindOne(bson.M{"event": event, "user_id": session.UserId})
}

func ListTicketByUser(page int64, pageSize int64, sort string, session types.Session) (ListOutput, error) {
	filter := bson.M{"user_id": session.UserId}
	return List(filter, page, pageSize, sort)
}
