package main

import (
	"context" //**
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"           //**
	"go.mongodb.org/mongo-driver/bson/primitive" //**
	"go.mongodb.org/mongo-driver/mongo"          //**
	"go.mongodb.org/mongo-driver/mongo/options"  //**
)

// Config
var mongoUri string = "mongodb://localhost:27017"
var mongoDbName string = "ars_app_db"

// Database variables
var mongoClient *mongo.Client
var flightCollection *mongo.Collection

type Flight struct {
	Id          string  `json:"id,omitempty" bson:"_id,omitempty"`
	Number      string  `json:"number" bson:"number"`
	AirlineName string  `json:"airline_name" bson:"airline_name"`
	Source      string  `json:"source" bson:"source"`
	Destination string  `json:"destination" bson:"destination"`
	Capacity    int     `json:"capacity" bson:"capacity"`
	Price       float64 `json:"price" bson:"price"`
}

// mongo connect
func connectToMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Println("Mongo DB Connection Error!" + err.Error())
		return
	}
	name := "flights"
	flightCollection = mongoClient.Database(mongoDbName).Collection(name)
}

// apis
func createFlight(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//json body as data
	var jbodyFlight Flight

	if err := c.BindJSON(&jbodyFlight); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//**add flight to mongo

	result, err := flightCollection.InsertOne(ctx, jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	//**get id of created flight
	flightId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\nId Error."})
		return
	}
	//**query created Flight
	var createdFlight Flight
	err = flightCollection.FindOne(ctx, bson.M{"_id": flightId}).Decode(&createdFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	/*createdFlight := Flight{Id: "001", Number: "AI 845", AirlineName: "Air India",
	Source: "Mumbai", Destination: "Abu dhabi", Capacity: 300, Price: 5000.0}*/
	//response
	c.JSON(http.StatusCreated,
		gin.H{"message": "Flight Created Successfully", "flight": createdFlight})
}

func readAllFlights(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//**query all flights
	var flights []Flight
	cursor, err := flightCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	defer cursor.Close(ctx)
	//
	flights = []Flight{}
	err = cursor.All(ctx, &flights)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	/*flights := []Flight{{Id: "001", Number: "AI 845", AirlineName: "Air India",
		Source: "Mumbai", Destination: "Abu dhabi", Capacity: 300, Price: 5000.0},
		{Id: "002", Number: "6E 151	", AirlineName: "Indigo",
			Source: "Hydrabad", Destination: "Bangalore", Capacity: 300, Price: 1000.0},
	}*/
	//response
	c.JSON(http.StatusOK, flights)
}

func readFlightById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//
	id := c.Param("id")
	// Convert string ID to primitive.ObjectID
	flightId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//**query flight by id
	var flight Flight
	err = flightCollection.FindOne(ctx, bson.M{"_id": flightId}).Decode(&flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Flight Not Found."})
		return
	}
	//**if no flight send msg "not found"
	/*flight := Flight{Id: id, Number: "AI 845", AirlineName: "Air India",
	Source: "Mumbai", Destination: "Abu dhabi", Capacity: 300, Price: 5000.0}*/
	//fmt.Println(id)
	//response
	c.JSON(http.StatusOK, flight)
}

func updateFlight(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//
	id := c.Param("id")
	// Convert string ID to primitive.ObjectID
	flightId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//**query flight by id
	var oldFlight Flight
	err = flightCollection.FindOne(ctx, bson.M{"_id": flightId}).Decode(&oldFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Flight Not Found."})
		return
	}
	var jbodyFlight Flight
	err = c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	oldFlight.Price = jbodyFlight.Price
	//**update flight
	//var result *mongo.UpdateResult
	_, err = flightCollection.UpdateOne(ctx, bson.M{"_id": flightId}, bson.M{"$set": bson.M{"price":jbodyFlight.Price}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	/*if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight Not Found."})
		return
	}*/
	/*updatedFlight := Flight{Id: id, Number: "AI 845", AirlineName: "Air India",
	Source: "Mumbai", Destination: "Abu dhabi", Capacity: 300, Price: 5000.0}*/
	//fmt.Println(id)
	//response
	c.JSON(http.StatusOK, gin.H{"message": "Flight Updated Successfully", "flight": oldFlight})
}

func deleteFlight(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//
	id := c.Param("id")
	// Convert string ID to primitive.ObjectID
	flightId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//**query flight by id
	var oldFlight Flight
	err = flightCollection.FindOne(ctx, bson.M{"_id": flightId}).Decode(&oldFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Flight Not Found."})
		return
	}
	//delete
	_, err = flightCollection.DeleteOne(ctx, bson.M{"_id": flightId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully."})
}

func main() {
	//**connect to mongo<-DONE
	connectToMongo()
	//router
	r := gin.Default()
	//cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//routes
	r.POST("/flights", createFlight)
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightById)
	r.PUT("/flights/:id", updateFlight)
	r.DELETE("/flights/:id", deleteFlight)
	//server
	r.Run(":8080")
}

/*

package main
import "fmt"

func main() {
    var txt interface{}
    txt = "Hello World"
    greeting, ok := txt.(int)
    fmt.Println(greeting, ok)
}

*/
