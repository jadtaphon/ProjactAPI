package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"

)

func main() {
	e := echo.New()
	port := os.Getenv("PORT")
	uri := os.Getenv("MONGODB_URI")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:80"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	db, err := mgo.Dial(uri)
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &Handler{DB: db}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/getAll", h.getUser)
	e.GET("/getKey/:id", h.getKey)
	e.POST("/createqr", h.createqr)
	e.POST("/updateqr", h.updatekey)

	e.Logger.Fatal(e.Start(":" + port))
}

///////////database connecttion//////////////
// db, err := mongo.NewClient(options.Client().ApplyURI(url))
// if err != nil {
// 	log.Fatal(err)
// }
// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// err = db.Connect(ctx)
// if err != nil {
// 	log.Fatal(err)
// }
// defer db.Disconnect(ctx)
// err = client.Ping(ctx, readpref.Primary())
// if err != nil {
// 	log.Fatal(err)
// }

// quickstartDatabase := db.Database("test")
// testCollection := quickstartDatabase.Collection("qr_api")

// update, err := testCollection.InsertOne(ctx, bson.D{
// 	{Key: "status", Value: "true"},
// })
// log.Println(quickstartDatabase)
// log.Println(update.InsertedID)
// log.Println(db)

// err = client.Ping(db, readpref.Primary())
// if err != nil {
// 	log.Fatal(err)
// }
//////////////////////////////////////////////////////////////
