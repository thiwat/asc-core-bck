package ticket

import (
	"context"
	"math"
	"time"

	"asc-core/db"
	"asc-core/utils"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ticketCollection *mongo.Collection = db.GetCollection(
	"ticket",
	bson.D{
		{Key: "code", Value: 1},
	},
)

func findOne(filter bson.M) (Ticket, error) {
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

func create(ticket Ticket) (Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticket.Code = uuid.NewString()

	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()
	_, err := ticketCollection.InsertOne(ctx, ticket)
	if err != nil {
		return ticket, err
	}
	return findOne(bson.M{"code": ticket.Code})
}

func updateOne(filter bson.M, ticket Ticket) (Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticket.UpdatedAt = time.Now()
	_, err := ticketCollection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": ticket},
	)

	if err != nil {
		return Ticket{}, err
	}
	return findOne(filter)
}

func list(filter bson.M, page int64, pageSize int64, sort string) (ListOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var events = make([]Ticket, 0)
	var res ListOutput

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
