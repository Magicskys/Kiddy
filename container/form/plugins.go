package form

import "gopkg.in/mgo.v2/bson"

type Plugins struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Pinyin string `json:"pinyin"`
	Title string `json:"title"`
	Classification string `json:"classification"`
	Danger string `json:"danger"`
}

type PostPlugin struct {
	Title string `form:"title" binding:"required"`
	Classification string `form:"classification" binding:"required"`
	Danger string `form:"danger" binding:"required"`
	Poc string `form:"poc" binding:"required"`
}

type PluginsPoc struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Pinyin string `json:"pinyin"`
	Title string `json:"title"`
	Classification string `json:"classification"`
	Danger string `json:"danger"`
	Poc string `json:"poc"`
}