package order

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

var orderCollection *mongo.Collection = db.GetCollection(
	"orderCollection",
	bson.D{
		{Key: "order_no", Value: 1},
	},
)

func findOne(filter bson.M) (Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var order Order

	err := orderCollection.FindOne(
		ctx,
		filter,
	).Decode(&order)

	if err != nil {
		return order, err
	}

	return order, nil
}

func create(order Order) (Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	order.OrderNo = uuid.NewString()

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	_, err := orderCollection.InsertOne(ctx, order)
	if err != nil {
		return order, err
	}
	return findOne(bson.M{"order_no": order.OrderNo})
}

func updateOne(filter bson.M, order Order) (Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	order.UpdatedAt = time.Now()
	_, err := orderCollection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": order},
	)

	if err != nil {
		return Order{}, err
	}
	return findOne(filter)
}

func list(filter bson.M, page int64, pageSize int64, sort string) (ListOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var events = make([]Order, 0)
	var res ListOutput

	sortObj := utils.BuildSort(sort)

	opts := options.Find()
	opts.SetSkip((page - 1) * pageSize)
	opts.SetLimit((pageSize))
	opts.SetSort(sortObj)

	cursor, err := orderCollection.Find(
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

	count, err := orderCollection.CountDocuments(ctx, filter)
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
