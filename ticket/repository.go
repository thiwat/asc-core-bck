package ticket

import (
	"context"
	"fmt"
	"math"
	"time"

	"asc-core/db"
	"asc-core/types"
	"asc-core/utils"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ticketCollection *mongo.Collection = db.GetCollection(
	"ticket",
	bson.D{
		{Key: "user_id", Value: 1},
		{Key: "event", Value: 1},
	},
)

func FindOne(filter bson.M) (Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var ticket Ticket

	err := ticketCollection.FindOne(
		ctx,
		filter,
	).Decode(&ticket)

	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

func Create(ticket Ticket) (Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticket.Code = uuid.NewString()

	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()
	_, err := ticketCollection.InsertOne(ctx, ticket)
	if err != nil {
		return ticket, err
	}
	return FindOne(bson.M{"code": ticket.Code})
}

func List(filter bson.M, page int64, pageSize int64, sort string) (types.ListOutput[Ticket], error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var events = make([]Ticket, 0)
	var res types.ListOutput[Ticket]

	sortObj := utils.BuildSort(sort)

	opts := options.Find()
	opts.SetSkip((page - 1) * pageSize)
	opts.SetLimit((pageSize))
	opts.SetSort(sortObj)

	cursor, err := ticketCollection.Find(
		ctx,
		filter,
		opts,
	)

	if err != nil {
		return res, err
	}

	if err = cursor.All(ctx, &events); err != nil {
		return res, err
	}

	count, err := ticketCollection.CountDocuments(ctx, filter)
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