package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

)

func main() {

	port := os.Getenv("PORT")
	e := echo.New()

	url := "mongodb+srv://jadtaphon:hbrY7322@cluster0-vkyg7.mongodb.net/test?retryWrites=true&w=majority"

	h := &Handler{URL: url}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/createqr", h.createqr)
	// e.GET("check_key/:id", h.checkkey)
	e.GET("/gettest/:id", h.getUser)
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
