package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"gopkg.in/mgo.v2/bson"

)

func (h *Handler) getUser(c echo.Context) (err error) {
	// users := []*DataQR{}

	id := c.Param("id")
	db, err := mongo.NewClient(options.Client().ApplyURI(h.URL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(ctx)

	objID, _ := primitive.ObjectIDFromHex(id)
	cursor, err := db.Database("test").Collection("qr_api").Find(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
	}
	var qrury []bson.M
	if err = cursor.All(ctx, &qrury); err != nil {
		log.Fatal(err)
	}
	log.Println(id)

	// defer db.Close()
	return c.JSON(http.StatusOK, qrury)
}

func (h *Handler) createqr(c echo.Context) (err error) {

	status := c.FormValue("status")
	idcoure := c.FormValue("idcoure")

	db, err := mongo.NewClient(options.Client().ApplyURI(h.URL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(ctx)

	update, err := db.Database("test").Collection("qr_api").InsertOne(ctx, bson.D{
		{Key: "status", Value: status},
		{Key: "idcoure", Value: idcoure},
	})

	return c.JSON(http.StatusOK, update)
}

// // func (u *Handler) checkkey(c echo.Context) (err error)  {
// // 	key:= c.Param("id")
// // 	keyqr, err := strconv.Atoi(key)
// // 	// if u.Data.Key==keyqr {
// // 	// 	return c.JSON(http.StatusOK, u.Status)
// // 	// }
// // 	log.Println(key)
// // 	log.Println(keyqr)
// // 		return c.JSON(http.StatusOK, u)
// // }
