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
	//"go.mongodb.org/mongo-driver/mongo/readpref"

)

func (h *Handler) getALL(c echo.Context) (err error) {
	db, err := mongo.NewClient(options.Client().ApplyURI(h.URL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 25*time.Second)
	err = db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(ctx)

	qrury, err := db.Database("testAPL").Collection("qr_api").Find(ctx, bson.M{})
	var result []bson.M
	if err = qrury.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	
	return c.JSON(http.StatusOK, result)
}

//////////////////////////////////////////////////////////////////////////////////////
func (h *Handler) getUser(c echo.Context) (err error) {

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
	qrury, err := db.Database("testAPL").Collection("qr_api").Find(ctx, bson.M{"key": objID})
	if err != nil {
		log.Fatal(err)
	}
	var result []bson.M
	if err = qrury.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, result)
}

//////////////////////////////////////////////////////////////////////////////////////
func (h *Handler) createqr(c echo.Context) (err error) {

	id := primitive.NewObjectID()
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

	update, err := db.Database("testAPL").Collection("qr_api").InsertOne(ctx, bson.D{
		{Key: "key", Value: id},
		{Key: "idcoure", Value: idcoure},
	})

	log.Println(update)

	return c.JSON(http.StatusOK, id)
}

//////////////////////////////////////////////////////////////////////////////////////
func (h *Handler) upadtestatus(c echo.Context) (err error) {

	idcoure := c.FormValue("idcoure")
	newKey := primitive.NewObjectID()
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

	result, err := db.Database("testAPL").Collection("qr_api").UpdateOne(
		ctx,
		bson.M{"idcoure": idcoure},
		bson.D{
			{"$set", bson.D{{"key", newKey}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	// log.Println(idcoure)
	// log.Printf("Updated %v Documents!\n", result.ModifiedCount)

	return c.JSON(http.StatusOK, idcoure)
}

//5e3ef573f12e89f890c4b248

//////////////////////////////////////////////////////////////////////////////////////
// func (u *Handler) checkkey(c echo.Context) (err error)  {
// 	key:= c.Param("id")
// 	keyqr, err := strconv.Atoi(key)
// 	// if u.Data.Key==keyqr {
// 	// 	return c.JSON(http.StatusOK, u.Status)
// 	// }
// 	log.Println(key)
// 	log.Println(keyqr)
// 		return c.JSON(http.StatusOK, u)
// }
