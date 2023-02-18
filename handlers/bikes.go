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
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	if err = cur.All(context.TODO(), &bikes); err != nil {
		return err
	}
	return c.JSON(bikes)
}

func hasUserAlreadyRented(sessionId string, coll *mongo.Collection) error {
	filter := bson.D{{Key: "session_id", Value: sessionId}}
	var result Bike
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}
	return errors.New("have already rented another bike")
}

func doesUserOwnsTheBike(bikePayload *Bike, coll *mongo.Collection) error {
	filter := bson.D{{Key: "_id", Value: bikePayload.ID}}
	var result Bike
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return err
	}
	if result.SessionId != bikePayload.SessionId {
		return errors.New("cannot return someone else's bike")
	}
	return nil
}

func UpdateBike(c *fiber.Ctx) error {
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}
	bikePayload := new(Bike)

	if err = c.BodyParser(&bikePayload); err != nil {
		return err
	}

	coll := client.Database(db.Database).Collection(db.BikesCollection)

	if bikePayload.Rented { // user wants to rent
		if err = hasUserAlreadyRented(bikePayload.SessionId, coll); err != nil {
			return err
		}
	}

	if !bikePayload.Rented { // user wants to return
		if err = doesUserOwnsTheBike(bikePayload, coll); err != nil {
			return err
		}
	}

	filter := bson.D{{Key: "_id", Value: bikePayload.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "session_id", Value: bikePayload.SessionId}, {Key: "rented", Value: bikePayload.Rented}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	return c.JSON(result)
}
