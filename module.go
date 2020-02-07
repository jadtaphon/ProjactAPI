package main

import (
	"gopkg.in/mgo.v2/bson"

)

type (
	DataQR struct {
		ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Datas Data          `json:"data" bson:"data"`
	}
	Data struct {
		Key    int  `json:"key" bson:"key"`
		Status bool `json:"status" bson:"status"`
	}
)


