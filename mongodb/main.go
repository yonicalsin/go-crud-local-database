package main

import (
	"context"
	"log"

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

		instance = &DatabaseFactory{conn: client.Database("first-example-with-mongodb")}
	}
	return instance
}

func ModelUser() {
	model := connect().conn.Collection("users")
}

func getAllUsers(c echo.Context) error {
}

func main() {

	// Server
	server := echo.New()

	// Routes
	server.GET("/users", getAllUsers)

	// Echo Server initializing
	server.Logger.Fatal(server.Start(":4000"))

}
