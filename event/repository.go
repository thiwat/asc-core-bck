package event

import (
	"context"
	"time"
	"math"

	"asc-core/db"
	"asc-core/types"
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

func List(page int64, pageSize int64, sort string) (types.ListOutput[Event], error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var events = make([]Event, 0)
	var res types.ListOutput[Event]

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