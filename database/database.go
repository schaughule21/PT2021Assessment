package database

import (
	"31Aug-Assessment/helpers"
	"31Aug-Assessment/models"
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateConnection(url string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	clientOptions := options.Client().ApplyURI(url)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, cancel, err
}

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func CreateFactory(factory models.Factory) error {
	client, context, cancel, err := CreateConnection("mongodb://localhost:27017/")
	CheckError(err)
	defer Close(client, context, cancel)
	collection := client.Database("SugarFactory").Collection("factory")
	factory.SrNo = helpers.Uuid(1)
	factory.RegiTime = primitive.NewDateTimeFromTime(time.Now())
	validate := validator.New()
	err = validate.Struct(factory)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context, factory)
	CheckError(err)
	return nil
}

func GetFactory() []models.Factory {
	client, context, cancel, err := CreateConnection("mongodb://localhost:27017/")
	CheckError(err)
	defer Close(client, context, cancel)
	collection := client.Database("SugarFactory").Collection("factory")
	res, err := collection.Find(context, bson.D{})
	CheckError(err)
	var factory []models.Factory
	if err = res.All(context, &factory); err != nil {
		log.Fatal(err)
	}
	return factory
}

func CheckValidation(user models.Factory) models.Factory {
	client, context, cancel, err := CreateConnection("mongodb://localhost:27017/")
	CheckError(err)
	defer Close(client, context, cancel)
	collection := client.Database("SugarFactory").Collection("factory")
	var res models.Factory
	err = collection.FindOne(context, bson.D{primitive.E{Key: "uname", Value: user.UName}}).Decode(&res)
	CheckError(err)
	return res
}

func CheckError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
