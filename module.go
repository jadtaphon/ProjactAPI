package main

import (
	"gopkg.in/mgo.v2/bson"

)

type (
	 DataQR struct {
		ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Datas []Data  			 `json:"data" bson:"data"`
	}
	Data struct{
		Key   string     	`json:"key" bson:"key"`  
		Status string  	`json:"status" bson:"status"`
		
	}
)


