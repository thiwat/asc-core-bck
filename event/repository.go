package event

import (
	"context"
	"math"
	"time"

	"asc-core/db"
	"asc-core/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var eventCollection *mongo.Collection = db.GetCollection(
	"event",
	bson.D{
		{Key: "code", Value: 1},
	},
)

func findOne(filter bson.M) (Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var event Event

	err := eventCollection.FindOne(
		ctx,
		filter,
	).Decode(&event)

	if err != nil {
		return event, err
	}

	return event, nil
}

func create(event Event) (Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()
	_, err := eventCollection.InsertOne(ctx, event)
	if err != nil {
		return event, err
	}
	return findOne(bson.M{"code": event.Code})
}

func updateOne(filter bson.M, event Event) (Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event.UpdatedAt = time.Now()
	_, err := eventCollection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": event},
	)

	if err != nil {
		return Event{}, err
	}
	return findOne(filter)
}

func list(page int64, pageSize int64, sort string) (ListOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var events = make([]Event, 0)
	var res ListOutput

	sortObj := utils.BuildSort(sort)

	opts := options.Find()
	opts.SetSkip((page - 1) * pageSize)
	opts.SetLimit((pageSize))
	opts.SetSort(sortObj)

	cursor, err := eventCollection.Find(
		ctx,
		bson.M{},
		opts,
	)

	if err != nil {
		return res, err
	}

	if err = cursor.All(ctx, &events); err != nil {
		return res, err
	}

	count, err := eventCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	res.Rows = events
	res.Page = page
	res.PageSize = pageSize
	res.Total = count
	res.TotalPage = int64(math.Ceil(float64(count) / float64(pageSize)))

	return res, nil
}
