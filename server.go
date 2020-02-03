
package main

import (
	"net/http"

	"github.com/labstack/echo"

)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/gettest",getUser)
	e.Logger.Fatal(e.Start(":80"))
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	id:="123456789 test"
	return c.JSON(http.StatusOK, id)
  //return c.String(http.StatusOK, id)
}
