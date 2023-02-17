package handlers

import (
	"context"

	"github.com/bike-sharing-app/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
