package main

import (
	"gopkg.in/mgo.v2/bson"

)

type (
	DataQR struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		CourseID  string        `json:"course_id" bson:"course_id"`
		CourseKEY string        `json:"course_key" bson:"course_key"`
		Key       bson.ObjectId `json:"key" bson:"key"`
		Url       string        `json:"url" bson:"url"`
	}
)
