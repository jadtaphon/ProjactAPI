package main

import (
	//"math/rand"
	//"bytes"
	"net/http"
	//"io/ioutil"
	//"strconv"
	"log"

	"github.com/labstack/echo/v4"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

)
func (h *Handler) getUser(c echo.Context) (err error) {
	users := []*DataQR{}
	db := h.DB.Clone()

	if err = db.DB("test").C("data_T").Find(nil).All(&users); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer db.Close()
	return c.JSON(http.StatusOK, &users)
}

func (h *Handler) createqr(c echo.Context) (err error) {	
	key := c.FormValue("key")
	status := c.FormValue("status")

// 	var bodyBytes []byte
// 	if c.Request().Body != nil {
// 		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
// 	}
//  //  id := c.Param("id")
   

// 	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	log.Print("hello")
	user := new(DataQR)
	update := Data{
		Key:key,
		Status:status,
	}
		
	
	user.Datas=append(user.Datas, update)
	log.Println(user.Datas)
	// err = c.Bind(user)
	// if err != nil {
	// 	log.Print(err)
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("test").C("data_T").Insert(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, &user)

}
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