package event

import (
	"go.mongodb.org/mongo-driver/bson"
)

func GetEvent(code string) (Event, error) {
	return findOne(bson.M{"code": code})
}

func CreateEvent(event Event) (Event, error) {
	return create(event)
}

func UpdateByCode(code string, event Event) (Event, error) {
	return updateOne(bson.M{"code": code}, event)
}

func ListEvent(page int64, pageSize int64, sort string) (ListOutput, error) {
	return list(page, pageSize, sort)
}
