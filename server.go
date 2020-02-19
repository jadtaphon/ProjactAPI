package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"

)

func main() {

	port := os.Getenv("PORT")
	//uri := os.Getenv("mongodb://heroku_4v7cvj1l:1uvhbepo2nqnspk8l2jsmpvugf@ds043942.mlab.com:43942/heroku_4v7cvj1l")

	e := echo.New()
	db, err := mgo.Dial("localhost")
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
	// e.GET("check_key/:id", h.checkkey)

	e.Logger.Fatal(e.Start(":" + port))
}
