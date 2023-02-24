package handlers

import (
	"context"
	"errors"

	"github.com/bike-sharing-app/db"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bike struct {
	ID        primitive.ObjectID `json:"id" bson:"_id" validate:"required"`
	Name      string             `json:"name" bson:"name"`
	Latitude  float64            `json:"latitude" bson:"latitude"`
	Longitude float64            `json:"longitude" bson:"longitude"`
	Rented    bool               `json:"rented" bson:"rented"`
	SessionId string             `json:"sessionId" bson:"session_id" validate:"required"`
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
	if len(bikes) == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Bikes not found!")
	}
	return c.JSON(bikes)
}

func getBike(bikeId primitive.ObjectID, coll *mongo.Collection) (Bike, error) {
	filter := bson.D{{Key: "_id", Value: bikeId}}
	var result Bike
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

func doesBikeExists(bikeId primitive.ObjectID, coll *mongo.Collection) error {
	if _, err := getBike(bikeId, coll); err != nil {
		return err
	}
	return nil
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
	if result, err := getBike(bikePayload.ID, coll); err != nil {
		return err
	} else {
		if result.SessionId != bikePayload.SessionId {
			return errors.New("cannot return someone else's bike")
		}
	}
	return nil
}

func isBikeRented(bikeId primitive.ObjectID, coll *mongo.Collection) error {
	if result, err := getBike(bikeId, coll); err != nil {
		return err
	} else {
		if !result.Rented {
			return errors.New("bike not rented, please rent before returning")
		}
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

	validate := validator.New()
	if err = validate.Struct(bikePayload); err != nil {
		return err
	}

	coll := client.Database(db.Database).Collection(db.BikesCollection)

	if err = doesBikeExists(bikePayload.ID, coll); err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Bike not found!")
	}

	rentStatus := bikePayload.Rented
	sessionId := ""

	if rentStatus { // user wants to rent
		sessionId = bikePayload.SessionId
		if err = hasUserAlreadyRented(sessionId, coll); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	if !rentStatus { // user wants to return
		if err = isBikeRented(bikePayload.ID, coll); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		if err = doesUserOwnsTheBike(bikePayload, coll); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	filter := bson.D{{Key: "_id", Value: bikePayload.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "session_id", Value: sessionId}, {Key: "rented", Value: rentStatus}}}}
	if _, err = coll.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString("Bike's rent status updated successfully!")
}
