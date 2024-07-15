package event

import (
	"context"
	"time"

	"asc-core/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var eventCollection *mongo.Collection = db.GetCollection(
	"event",
	bson.D{
		{Key: "code", Value: 1},
	},
)

func FindByCode(code string) (Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var event Event

	err := eventCollection.FindOne(
		ctx,
		bson.M{"code": code},
	).Decode(&event)

	if err != nil {
		return event, err
	}

	return event, nil
}

func Create(event Event) (Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()
	_, err := eventCollection.InsertOne(ctx, event)
	if err != nil {
		return event, err
	}
	return FindByCode(event.Code)
}
