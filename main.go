package main

import (
	"blog-platform/resource"
	"blog-platform/router"
	"context"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Init echo framework
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init a empty context
	ctx := context.Background()

	InitLog(e)

	InitMongoClient(ctx)
	// Close mongo client connection when the program exits
	defer func() {
		if err := resource.MongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// validator
	resource.Validate = validator.New(validator.WithRequiredStructEnabled())

	// Register routes
	router.RegisterRoutes(e)

	// Get port
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		// For debugging locally
		httpPort = "8082"
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + httpPort))

}

// Init log
func InitLog(e *echo.Echo) {
	// For echo
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		e.Logger.Fatal(err)
		panic(err.Error())
	}
	e.Logger.SetOutput(logFile)

	// e.Debug = true
}

// Init mongo client
func InitMongoClient(ctx context.Context) {
	var err error

	// Set up mongo connection information
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")

	// Connect to MongoDB
	resource.MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = resource.MongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

}
