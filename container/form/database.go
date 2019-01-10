package form

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Info struct {
	ID bson.ObjectId `json:"uid" bson:"_id"`
	Method string `json:"method"`
	Host string `json:"host"`
	URL string `json:"url"`
	Cookies string `json:"cookies"`
	Body string `json:"body"`
	Scheme string `json:"scheme"`
}

type GeneralResult struct {
	ID bson.ObjectId `json:"taskId" bson:"_id"`
	Datetime time.Time	`json:"datetime"`
	Class string `json:"class"`
	Typeof string `json:"typeof"`
	Host string `json:"host"`
	Message string `json:"message"`
	Status string `json:"status"`
}