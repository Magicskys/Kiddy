package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

func Insert_Sqlmap_Task(name string,objectId string,taskId string) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.Insert(bson.M{"class":"sqlinject","taskId":taskId,"objectId":bson.ObjectIdHex(objectId),"datetime":time.Now()})
	if err!=nil{
		return false
	}
	return true
}

func remove_sqlmap_result(taskId string) bool {
	mgoc,session:=CopySession("sqlmap")
	defer session.Close()
	err:=mgoc.Remove(bson.M{"taskId":taskId})
	if err!=nil{
		return false
	}
	return true
}

func Remmove_name_taskId(name string,taskId string) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.Remove(bson.M{"taskId":taskId})
	if err!=nil{
		return false
	}
	if !remove_sqlmap_result(taskId){
		return false
	}
	return true
}
