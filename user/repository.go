package user

import (
	"asc-core/db"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.GetCollection(
	"user",
	bson.D{
		{Key: "user_id", Value: 1},
	},
)

func findOne(filter bson.M) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user User

	err := userCollection.FindOne(
		ctx,
		filter,
	).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func Create(user User) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}
	return findOne(bson.M{"user_id": user.UserId})
}

func UpdateOne(filter bson.M, user User) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.UpdatedAt = time.Now()

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": user},
	)

	if err != nil {
		return user, err
	}
	return findOne((filter))
}
