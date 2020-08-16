package main

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User is ...
type User struct {
	ID       string `bson:"-"`
	Name     string `bson:"name"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

// DatabaseFactory is ...
type DatabaseFactory struct {
	conn *mongo.Database
	user *mongo.Collection
}

var instance *DatabaseFactory

func connect() *DatabaseFactory {
	if instance == nil {
		// Database Connect
		clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(context.TODO(), clientOpts)
		if err != nil {
			log.Fatal(err)
		}

		// Check the connections
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}

		conn := client.Database("first-example-with-mongodb")

		// Collections
		instance = &DatabaseFactory{
			conn: conn,
			user: conn.Collection("User"),
		}

	}
	return instance
}

func getAllUsers(ctx echo.Context) error {
	options := options.Find()
	options.SetLimit(3)

	var results []*User

	cur, err := connect().user.Find(context.TODO(), bson.D{{}}, options)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s User
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return ctx.JSONPretty(http.StatusOK, results, "  ")
}

func main() {

	// Server
	server := echo.New()

	// Routes
	server.GET("/users", getAllUsers)

	// Echo Server initializing
	server.Logger.Fatal(server.Start(":4000"))

}
