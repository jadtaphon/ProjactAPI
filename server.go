
package main

import (
	"net/http"
	"os"

	//"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"

)

func main() {

	port:= os.Getenv("PORT")
	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))

	// db, err := mgo.Dial("localhost")

	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// hub := newHub()
	// go hub.run()

	h := &Handler{DB: db}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/createqr/", h.createqr)
	// e.GET("check_key/:id", h.checkkey)
	e.GET("/gettest",h.getUser)
	e.Logger.Fatal(e.Start(":"+port))
}


