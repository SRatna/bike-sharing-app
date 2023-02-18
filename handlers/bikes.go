package handlers

import (
	"context"
	"errors"

	"github.com/bike-sharing-app/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bike struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Latitude  float64            `json:"latitude" bson:"latitude"`
	Longitude float64            `json:"longitude" bson:"longitude"`
	Rented    bool               `json:"rented" bson:"rented"`
	SessionId string             `json:"sessionId" bson:"session_id"`
}

func GetAllBikes(c *fiber.Ctx) error {
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}
	var bikes []Bike
	coll := client.Database(db.Database).Collection(db.BikesCollection)
	cur, err := coll.Find(context.TODO(), bson.D{
		primitive.E{},
	})
	if err != nil {
		return err
	}
	if err = cur.All(context.TODO(), &bikes); err != nil {
		return err
	}
	return c.JSON(bikes)
}

func hasUserAlreadyRented(bike *Bike, coll *mongo.Collection) error {
	filter := bson.D{{Key: "session_id", Value: bike.SessionId}}
	var result Bike
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}
	return errors.New("You have already rented another bike.")
}

func UpdateBike(c *fiber.Ctx) error {
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}
	bike := new(Bike)

	if err = c.BodyParser(&bike); err != nil {
		return err
	}

	coll := client.Database(db.Database).Collection(db.BikesCollection)

	if err = hasUserAlreadyRented(bike, coll); err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: bike.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "session_id", Value: bike.SessionId}, {Key: "rented", Value: bike.Rented}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	return c.JSON(result)
}
