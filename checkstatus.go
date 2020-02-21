package main

import (
	//"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"

)

//////////////////////////////////////////////////////////////////////////////////////////
func (h *Handler) getUser(c echo.Context) (err error) {
	users := []*DataQR{}
	db := h.DB.Clone()

	if err = db.DB("heroku_4v7cvj1l").C("qr_api").Find(nil).All(&users); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer db.Close()
	return c.JSON(http.StatusOK, &users)
}

////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////getKey///////////////////////////////////////////////////////////
func (h *Handler) getKey(c echo.Context) (err error) {
	users := new(DataQR)
	if err = c.Bind(users); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()

	id := c.Param("id")

	if err = db.DB("heroku_4v7cvj1l").C("qr_api").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&users); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	//log.Println(users.Url)
	//return c.JSON(http.StatusOK, users.Url)
	return c.Redirect(http.StatusMovedPermanently, users.Url)
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////createqr////////////////////////////////////////////////////////////
func (h *Handler) createqr(c echo.Context) (err error) {

	user := &DataQR{ID: bson.NewObjectId()}
	err = c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("heroku_4v7cvj1l").C("qr_api").Insert(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}

///////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////updatekey/////////////////////////////////////////////////////////////
func (h *Handler) updatekey(c echo.Context) (err error) {
	//users := []*DataQR{}
	users := new(DataQR)
	if err = c.Bind(users); err != nil {
		return
	}
	courseid := c.FormValue("courseid")
	coursekey := c.FormValue("coursekey")
	urls := c.FormValue("url")

	db := h.DB.Clone()
	defer db.Close()

	query := bson.M{"course_id": courseid}
	update := bson.M{"$set": bson.M{"course_key": coursekey, "url": urls}}

	if err = db.DB("heroku_4v7cvj1l").C("qr_api").Update(query, update); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
